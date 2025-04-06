package com.github.walteh.cloudstack.proxy.generator

import org.apache.cloudstack.api.BaseCmd.CommandType
import com.cloud.serializer.Param
import org.apache.cloudstack.acl.RoleType
import org.apache.cloudstack.acl.SecurityChecker
import org.apache.cloudstack.api.*
import java.lang.reflect.Field
import kotlin.reflect.jvm.jvmName

// Data classes for metadata representation
data class GroupMetadata(
    val name: String,
    val scope: String,
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
	val commandType: BaseCmd.CommandType,
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
	val authorized: List<RoleType> = emptyList(),
	val collectionType: BaseCmd.CommandType,
	val length: Int,

//		x = parameter.acceptedOnAdminPort,
//	y = parameter.expose,
//	z = parameter.includeInApiDoc,
//	a = parameter.authorized,
//	b = parameter.collectionType,
//	c = parameter.length,
//	d = parameter.validations,


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
	val accessType: SecurityChecker.AccessType,
//	val roleTypes: List<String>
)

data class AuthorizationInfo(
    val roleTypes: List<String>
)

/**
 * Extract metadata from command classes and write to JSON files
 */
fun extractCommandsToJson(commandClasses: List<Class<*>>): List<GroupMetadata?>  {
    println("Extracting metadata from ${commandClasses.size} command classes")
    
	val commandMetadata = commandClasses.mapNotNull { extractCommandMetadata(it) }
    // Group commands by API group
    val commandsByGroup = groupCommandsByApi(commandMetadata)
    
    // Process each group
	return commandsByGroup.map { (groupName, commands) ->
        if (commands.isNotEmpty()) {
            // Group commands by scope (admin, user, etc.)
            val scopedCommands = commands.groupBy { it.scopeName }
            
            // Process each scope group
            scopedCommands.map { (scope, cmds) ->
                // Collect all response objects used by commands in this group
                val responseObjectClasses = collectResponseObjectClasses(cmds)
                val responseObjectMetadata = responseObjectClasses.mapNotNull { extractResponseObjectMetadata(it) }
                
                // Collect all entity types referenced by commands in this group
                val entityTypes = collectEntityTypes(cmds)
                
                GroupMetadata(
                    name = groupName,
                    scope = scope,
                    commands = cmds,
                    responseObjects = responseObjectMetadata,
                    entityTypes = entityTypes
                )
            }.firstOrNull() // Return first scope group or null
        } else {
            null
        }
    }
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
    val entityReference = responseClass.getAnnotation(EntityReference::class.java)?.value?.first()?.java?.name
    
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
        .map { extractParameterMetadata(it) }
    
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
        accessType = aclAnnotation.accessType,
//        roleTypes = aclAnnotation.map { it.name }
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
            .map { extractParameterMetadata(it) }
        
        inheritedParameters.addAll(fields)
        currentClass = currentClass.superclass
    }
    
    return inheritedParameters
}

/**
 * Extract metadata from a parameter field
 */
fun extractParameterMetadata(field: Field): ParameterMetadata {
    val parameter = field.getAnnotation(Parameter::class.java)
    
    // Get validation annotations
    val validations = extractValidations(field)
    
    // Extract authorization info if present
    val authorizationInfo = extractAuthorizationInfo(parameter)
    
    return ParameterMetadata(
        name = parameter.name,
        fieldName = field.name,
        javaType = field.type.name,
        commandType = parameter.type,
        description = parameter.description,
        required = parameter.required,
        since = parameter.since,
		acceptedOnAdminPort = parameter.acceptedOnAdminPort,
		expose = parameter.expose,
		includeInApiDoc = parameter.includeInApiDoc,
		authorized = parameter.authorized.toList(),
		collectionType = parameter.collectionType,
		length = parameter.length,
        entityType = parameter.entityType.firstOrNull()?.javaObjectType?.typeName,
        validations = validations,
        authorization = authorizationInfo,
        shouldDisplay = true
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