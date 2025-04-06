package com.github.walteh.cloudstack.proxy.generator

import org.apache.cloudstack.api.BaseCmd.CommandType
import com.cloud.serializer.Param
import org.apache.cloudstack.acl.RoleType
import org.apache.cloudstack.acl.SecurityChecker.AccessType
import org.apache.cloudstack.api.*
import java.lang.reflect.Field
import java.lang.reflect.Method
import kotlin.reflect.jvm.jvmName

// Data classes for metadata representation
data class GroupMetadata(
    val name: String,
    val commands: List<CommandMetadata>,
    val responseObjects: List<ResponseObjectMetadata> = emptyList(),
    val entityTypes: List<String> = emptyList()
)

data class CommandMetadata(
	val groupName: String,
	val scopeName: String,
	val scopedReplicaOf: String,
	val className: String,
	val simpleName: String,
	val commandName: String,
	val description: String,
	val responseObject: String,
	val responseView: String,
	val entityType: List<String>,
	val parameters: List<ParameterMetadata>,
	val superclasses: List<String> = emptyList(),
	val isAsync: Boolean = false,
	val aclInfo: ACLInfo? = null,
	val requestHasSensitiveInfo: Boolean = false,
	val responseHasSensitiveInfo: Boolean = false,
)

data class ParameterMetadata(
	val name: String,
	val fieldName: String,
	val javaType: String,
	val commandType: CommandType,
	val description: String,
	val required: Boolean,
	val since: String,
	val entityType: String? = null,
	val validations: List<ValidationMetadata> = emptyList(),
	val authorization: AuthorizationInfo? = null,
	val shouldDisplay: Boolean = true,
	val acceptedOnAdminPort: Boolean = false,
	val expose: Boolean = false,
	val includeInApiDoc: Boolean = false,
	val authorized: List<String> = emptyList(),
	val collectionType: CommandType? = null,
	val length: Int = 0,
	val getterMethod: GetterMethodInfo? = null,
	val isCommaSeparatedList: Boolean = false
)

data class GetterMethodInfo(
    val methodName: String,
    val returnType: String,
    val returnTypeEnum: Boolean,
    val isList: Boolean = false
)

data class ResponseObjectMetadata(
    val className: String,
    val simpleName: String,
    val fields: List<ResponseFieldMetadata>,
    val entityReference: String? = null,
    val superclass: String? = null
)

data class ResponseFieldMetadata(
    val name: String,
    val fieldName: String,
    val javaType: String,
    val description: String,
    val since: String? = null,
    val authorized: List<String>? = null,
    val responseObject: String? = null
)

data class ValidationMetadata(
    val type: String,
    val value: String
)

data class ACLInfo(
	val accessType: AccessType
	// roleTypes removed as it doesn't exist
)

data class AuthorizationInfo(
    val roleTypes: List<String>
)

/**
 * Extract metadata from command classes and write to JSON files
 */
fun extractCommandsToMetadata(commandClasses: List<Class<*>>): Map<String, List<GroupMetadata>>  {
    println("Extracting metadata from ${commandClasses.size} command classes")
    
	val commandMetadata = commandClasses.mapNotNull { extractCommandMetadata(it) }
    // Group commands by API group
    val commandsByGroup = groupCommandsByApi(commandMetadata)
    
    // Process each group
	val metadataByGroup = commandsByGroup.map { (groupName, commands) ->
        if (commands.isNotEmpty()) {
            // Group commands by scope (admin, user, etc.)
			val responseObjectClasses = collectResponseObjectClasses(commands)
			val responseObjectMetadata = responseObjectClasses.mapNotNull { extractResponseObjectMetadata(it) }
			val entityTypes = collectEntityTypes(commands)

			GroupMetadata(
				name = groupName,
				commands = commands,
				responseObjects = responseObjectMetadata,
				entityTypes = entityTypes
			)

        } else {
            null
        }
    }

	return metadataByGroup.mapNotNull { it }.groupBy { it.name }
}

/**
 * Group command classes by API area
 */
fun groupCommandsByApi(commandClasses: List<CommandMetadata>): Map<String, List<CommandMetadata>> {
    return commandClasses.groupBy { clazz -> clazz.groupName }
}

/**
 * Collect all response object classes referenced by a list of commands
 */
fun collectResponseObjectClasses(commands: List<CommandMetadata>): List<Class<*>> {
    val responseObjectNames = commands.map { it.responseObject }.distinct()
    return responseObjectNames.mapNotNull { responseClassName ->
        try {
            Class.forName(responseClassName)
        } catch (e: Exception) {
            println("Could not load response class: $responseClassName - ${e.message}")
            null
        }
    }
}

/**
 * Collect all entity types referenced by commands
 */
fun collectEntityTypes(commands: List<CommandMetadata>): List<String> {
    return commands.flatMap { it.entityType }.filter { it.isNotEmpty() }.distinct()
}

/**
 * Extract metadata from a response object class
 */
fun extractResponseObjectMetadata(responseClass: Class<*>): ResponseObjectMetadata? {
    if (!BaseResponse::class.java.isAssignableFrom(responseClass)) {
        return null
    }
    
    // Get entity reference if present
    val entityReference = responseClass.getAnnotation(EntityReference::class.java)?.value?.firstOrNull()?.java?.name
    
    // Get superclass
    val superclass = responseClass.superclass?.name
    
    // Extract fields with Param annotation
    val fields = responseClass.declaredFields
        .filter { it.isAnnotationPresent(Param::class.java) }
        .map { field ->
            val param = field.getAnnotation(Param::class.java)
            ResponseFieldMetadata(
                name = findSerializedName(field),
                fieldName = field.name,
                javaType = field.type.name,
                description = param.description,
                since = param.since.takeIf { it.isNotEmpty() },
                authorized = extractAuthorizedRoles(param),
                responseObject = findResponseObject(param)
            )
        }
    
    return ResponseObjectMetadata(
        className = responseClass.name,
        simpleName = responseClass.simpleName,
        fields = fields,
        entityReference = entityReference,
        superclass = superclass
    )
}

/**
 * Find the serialized name of a field using SerializedName annotation
 */
fun findSerializedName(field: Field): String {
    val serializedNameAnn = field.annotations.find { it.annotationClass.java.simpleName == "SerializedName" }
    if (serializedNameAnn != null) {
        try {
            val value = serializedNameAnn.javaClass.getMethod("value").invoke(serializedNameAnn) as String
            return value
        } catch (e: Exception) {
            // Couldn't get value from annotation
        }
    }
    return field.name
}

/**
 * Extract authorized roles from a Param annotation
 */
fun extractAuthorizedRoles(param: Param): List<String>? {
    try {
        val authorizedMethod = param.javaClass.getMethod("authorized")
        val authorized = authorizedMethod.invoke(param) as Array<*>
        if (authorized.isNotEmpty()) {
            return authorized.map { it.toString() }
        }
    } catch (e: Exception) {
        // Authorized property doesn't exist or is not accessible
    }
    return null
}

/**
 * Find the response object class name from a Param annotation
 */
fun findResponseObject(param: Param): String? {
    try {
        val responseObjectMethod = param.javaClass.getMethod("responseObject")
        val responseObject = responseObjectMethod.invoke(param) as Class<*>
        if (responseObject != Void::class.java) {
            return responseObject.name
        }
    } catch (e: Exception) {
        // responseObject property doesn't exist or is not accessible
    }
    return null
}

/**
 * Extract metadata from a single command class
 */
fun extractCommandMetadata(commandClass: Class<*>): CommandMetadata? {
    val apiCommand = commandClass.getAnnotation(APICommand::class.java) ?: return null
    
    val parameters = commandClass.declaredFields
        .filter { it.isAnnotationPresent(Parameter::class.java) }
        .map { extractParameterMetadata(it, commandClass) }
    
    // Extract inherited command information
    val inheritedParameters = extractInheritedParameters(commandClass)
    val allParameters = parameters + inheritedParameters
    
    // Get additional metadata from annotations
    val aclInfo = extractACLInfo(commandClass)
    val isAsync = isAsyncCommand(commandClass)
    
    // Get inheritance hierarchy
    val superclasses = getSuperclasses(commandClass)

	val packageName = commandClass.packageName.removePrefix("org.apache.cloudstack.api.command.")
	var split = packageName.split(".")
	val userType = split.first()
	val rest = split.drop(1)
	val group = rest.joinToString(".")

	var scopedRep = ""

	if (userType == "admin" && split.last().endsWith("ByAdmin") && superclasses.first().equals(commandClass.packageName.replace(".admin.", ".user.").removeSuffix("ByAdmin"))) {
		scopedRep = superclasses.first()
	}

    return CommandMetadata(
		scopedReplicaOf = scopedRep,
		groupName = group,
		scopeName = userType,
        className = commandClass.name,
        simpleName = commandClass.simpleName,
        commandName = apiCommand.name,
        description = apiCommand.description,
        responseObject = apiCommand.responseObject.qualifiedName ?: "",
        responseView = apiCommand.responseView.name,
        entityType = apiCommand.entityType.map { it.qualifiedName ?: "" },
        parameters = allParameters,
        superclasses = superclasses,
        isAsync = isAsync,
        aclInfo = aclInfo,
        requestHasSensitiveInfo = apiCommand.requestHasSensitiveInfo,
        responseHasSensitiveInfo = apiCommand.responseHasSensitiveInfo
    )
}

/**
 * Extract ACL information from a command class
 */
fun extractACLInfo(commandClass: Class<*>): ACLInfo? {
    val aclAnnotation = commandClass.getAnnotation(ACL::class.java) ?: return null
    
    return ACLInfo(
        accessType = aclAnnotation.accessType
        // roleTypes removed as it doesn't exist
    )
}

/**
 * Check if a command is an async command
 */
fun isAsyncCommand(commandClass: Class<*>): Boolean {
    val asyncCommandBaseClasses = listOf(
        "BaseAsyncCmd",
        "BaseAsyncCreateCmd",
        "BaseAsyncCustomIdCmd"
    )
    
    return getSuperclasses(commandClass).any { superclass ->
        asyncCommandBaseClasses.any { baseClass -> 
            superclass.endsWith(baseClass) 
        }
    }
}

/**
 * Get the superclass hierarchy of a class
 */
fun getSuperclasses(clazz: Class<*>): List<String> {
    val superclasses = mutableListOf<String>()
    var currentClass: Class<*>? = clazz.superclass
    
    while (currentClass != null && currentClass != Object::class.java) {
        superclasses.add(currentClass.name)
        currentClass = currentClass.superclass
    }
    
    return superclasses
}

/**
 * Extract parameters from superclasses
 */
fun extractInheritedParameters(commandClass: Class<*>): List<ParameterMetadata> {
    val inheritedParameters = mutableListOf<ParameterMetadata>()
    var currentClass: Class<*>? = commandClass.superclass
    
    while (currentClass != null && currentClass != Object::class.java) {
        // Get declared fields with Parameter annotation
        val fields = currentClass.declaredFields
            .filter { it.isAnnotationPresent(Parameter::class.java) }
            .map { extractParameterMetadata(it, currentClass!!) }
        
        inheritedParameters.addAll(fields)
        currentClass = currentClass.superclass
    }
    
    return inheritedParameters
}

/**
 * Find a getter method for a parameter that provides stronger typing (String→Enum) or list handling
 */
fun findGetterMethodForParameter(parameterName: String, parameterField: Field, parameterType: Class<*>, clazz: Class<*>): GetterMethodInfo? {
    // Only look for special getters for String parameters (that could be enums) or list parameters
    if (parameterType != String::class.java && 
        parameterType != java.lang.String::class.java && 
        !parameterType.name.contains("List") &&
        !parameterType.isArray) {
        return null
    }
    
    // Convert parameter name to potential getter name
    val getterPrefix = "get"
    val capitalizedName = parameterName.replaceFirstChar { it.uppercaseChar() }
    val getterNames = listOf(
        "$getterPrefix$capitalizedName", 
        "$getterPrefix${parameterField.name.replaceFirstChar { it.uppercaseChar() }}",
        parameterName.camelToSnakeCase().camelCase(), // Handle special cases like API_KEY -> apiKey -> getApiKey
        // Try RoleType specific patterns
        "getRoleType", 
        "getRoleTypes",
        "getRole"
    )
    
    // Look for matching getter methods
    for (methodName in getterNames) {
        try {
            val methods = clazz.methods.filter { it.name.equals(methodName, ignoreCase = true) }
            for (method in methods) {
                val returnType = method.returnType
                val isEnum = returnType.isEnum
                val isList = returnType.name.contains("List") || returnType.isArray
                
                // Check if return type is an enum (for String params) or if it's a list-related getter
                if (isEnum || (isList && (parameterType == String::class.java || parameterType == java.lang.String::class.java))) {
                    return GetterMethodInfo(
                        methodName = method.name,
                        returnType = returnType.name,
                        returnTypeEnum = isEnum,
                        isList = isList
                    )
                }
                
                // Also look for methods that specifically handle comma-separated lists
                if (method.name.lowercase().contains("list") || 
                    method.returnType.simpleName.contains("List") ||
                    (parameterName.lowercase().contains("ids") && parameterType == String::class.java)) {
                    return GetterMethodInfo(
                        methodName = method.name,
                        returnType = returnType.name,
                        returnTypeEnum = isEnum,
                        isList = true
                    )
                }
                
                // Special check for RoleType
                if (method.name.lowercase().contains("roletype") || 
                    method.returnType.name.contains("RoleType")) {
                    return GetterMethodInfo(
                        methodName = method.name,
                        returnType = returnType.name,
                        returnTypeEnum = returnType.isEnum,
                        isList = returnType.isArray || returnType.name.contains("List")
                    )
                }
            }
        } catch (e: Exception) {
            // Method doesn't exist or can't be accessed
        }
    }
    
    // Check for common comma-separated list patterns in the parameter name or description
    val parameter = parameterField.getAnnotation(Parameter::class.java)
    if (parameter != null) {
        val desc = parameter.description.lowercase()
        if ((desc.contains("comma") && desc.contains("separat")) || 
            desc.contains("list of") ||
            parameterName.lowercase().contains("list") || 
            (parameterName.lowercase().endsWith("s") && 
             (parameterType == String::class.java || parameterType == java.lang.String::class.java))) {
            
            return GetterMethodInfo(
                methodName = "",  // No actual method
                returnType = "java.util.List",
                returnTypeEnum = false,
                isList = true
            )
        }
        
        // Check for role type in description
        if (desc.contains("role") && (desc.contains("type") || desc.contains("permission"))) {
            return GetterMethodInfo(
                methodName = "",
                returnType = "org.apache.cloudstack.acl.RoleType",
                returnTypeEnum = true,
                isList = desc.contains("list") || parameterName.lowercase().endsWith("s")
            )
        }
    }
    
    return null
}

/**
 * Convert camelCase to snake_case
 */
fun String.camelToSnakeCase(): String {
    return replace(Regex("([a-z])([A-Z]+)"), "$1_$2").lowercase()
}

/**
 * Convert snake_case to camelCase
 */
fun String.camelCase(): String {
    return split('_').mapIndexed { index, s ->
        if (index == 0) s.lowercase() else s.lowercase().capitalize()
    }.joinToString("")
}

/**
 * Extract metadata from a parameter field
 */
fun extractParameterMetadata(field: Field, commandClass: Class<*>): ParameterMetadata {
    val parameter = field.getAnnotation(Parameter::class.java)
    
    // Get validation annotations
    val validations = extractValidations(field)
    
    // Extract authorization info if present
    val authorizationInfo = extractAuthorizationInfo(parameter)
    
    // Find getter method that might provide stronger typing or list handling
    val getterInfo = findGetterMethodForParameter(parameter.name, field, field.type, commandClass)
    
    // Determine if this is likely a comma-separated list
    val isCommaSeparatedList = getterInfo?.isList == true || 
                              (parameter.description.lowercase().contains("comma") && 
                               parameter.description.lowercase().contains("separat")) ||
                              parameter.name.lowercase().contains("cidr") ||
                              (parameter.type == CommandType.LIST && field.type == String::class.java)
    
    // Get collection type with safe handling
    val collectionType = try {
        parameter.collectionType
    } catch (e: Exception) {
        null // Default value if not available
    }
    
    // Extract authorized roles with safe handling
    val authorizedRoles = try {
        parameter.authorized.map { it.name }
    } catch (e: Exception) {
        emptyList<String>()
    }
    
    return ParameterMetadata(
        name = parameter.name,
        fieldName = field.name,
        javaType = field.type.name,
        commandType = parameter.type,
        description = parameter.description,
        required = parameter.required,
        since = parameter.since,
        entityType = try { parameter.entityType.firstOrNull()?.qualifiedName } catch (e: Exception) { null },
        acceptedOnAdminPort = try { parameter.acceptedOnAdminPort } catch (e: Exception) { false },
        expose = try { parameter.expose } catch (e: Exception) { false },
        includeInApiDoc = try { parameter.includeInApiDoc } catch (e: Exception) { false },
        authorized = authorizedRoles,
        collectionType = collectionType,
        length = try { parameter.length } catch (e: Exception) { 0 },
        validations = validations,
        authorization = authorizationInfo,
        shouldDisplay = true,
        getterMethod = getterInfo,
        isCommaSeparatedList = isCommaSeparatedList
    )
}

/**
 * Extract authorization information from a parameter
 */
fun extractAuthorizationInfo(parameter: Parameter): AuthorizationInfo? {
    val authorized = parameter.authorized
    if (authorized.isEmpty()) {
        return null
    }
    
    return AuthorizationInfo(
        roleTypes = authorized.map { it.name }
    )
}

/**
 * Extract validation metadata from field annotations
 */
fun extractValidations(field: Field): List<ValidationMetadata> {
    val validations = mutableListOf<ValidationMetadata>()
    
    // Check for @Validate annotation
    if (field.isAnnotationPresent(Validate::class.java)) {
        val validate = field.getAnnotation(Validate::class.java)
        
        validations.add(
			ValidationMetadata(
            type = "validate",
            value = "true"  // This is simplified
        )
		)
    }
    
    // Add more validation annotations as needed
    
    return validations
} 