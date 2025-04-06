package com.github.walteh.cloudstack.proxy.generator

import org.apache.cloudstack.api.APICommand
import org.apache.cloudstack.api.Parameter
import org.apache.cloudstack.api.BaseCmd.CommandType
import org.apache.cloudstack.api.ACL
import org.apache.cloudstack.api.Validate
import java.lang.reflect.Field

// Data classes for metadata representation
data class GroupMetadata(
    val name: String,
    val commands: List<CommandMetadata>
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
	val commandType: String,
	val protoType: String,
	val protoTypeAnnotation: String? = null,
	val description: String,
	val required: Boolean,
	val since: String,
	val entityType: String? = null,
	val validations: List<ValidationMetadata> = emptyList(),
	val authorization: AuthorizationInfo? = null,
	val shouldDisplay: Boolean = true
)

data class ProtoTypeInfo(
    val type: String,
    val annotation: String? = null
)

data class ValidationMetadata(
    val type: String,
    val value: String
)

data class ACLInfo(
    val accessType: String,
    val roleTypes: List<String>
)

data class AuthorizationInfo(
    val roleTypes: List<String>
)

/**
 * Extract metadata from command classes and write to JSON files
 */
fun extractCommandsToJson(commandClasses: List<Class<*>>): List<GroupMetadata?>  {
    println("Extracting metadata from ${commandClasses.size} command classes")
    
    // Create output directory if it doesn't exist
//    outputDir.mkdirs()
	val commandMetadata = commandClasses.mapNotNull { extractCommandMetadata(it) }
    // Group commands by API group
    val commandsByGroup = groupCommandsByApi(commandMetadata)
    
    // Process each group
	return commandsByGroup.map { (groupName, commands) ->

        if (commandMetadata.isNotEmpty()) {
			GroupMetadata(
				name = groupName,
				commands = commands
			)


//            // Write to JSON file
//            val outputFile = File(outputDir, "$groupName.json")
//            val objectMapper = ObjectMapper().registerKotlinModule()
//            objectMapper.writerWithDefaultPrettyPrinter().writeValue(outputFile, groupMetadata)
//            println("Generated metadata for $groupName: ${outputFile.absolutePath}")
		} else {
			null
		}
    }


}

/**
 * Group command classes by API area
 */
fun groupCommandsByApi(commandClasses: List<CommandMetadata>): Map<String, List<CommandMetadata>> {
	// org.apache.cloudstack.api.command.admin.vpc.
    return commandClasses.groupBy { clazz ->
//        val packageName = clazz.packageName.removePrefix("org.apache.cloudstack.api.command.")
//		var split = packageName.split(".")
//		val userType = split.first()
//		val rest = split.drop(1)
//		val group = rest.joinToString(".")
//		group
		clazz.groupName
//        when {
//            packageName.contains("vpc") -> "vpc"
//            packageName.contains("vm") -> "vm"
//            packageName.contains("volume") -> "storage"
//            packageName.contains("network") -> "network"
//			packageName.contains("user") -> "user"
//            // Add more groups as needed
//            else -> "core"
//        }
    }
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
        entityType = apiCommand.entityType.map { it.simpleName ?: "" },
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
        accessType = aclAnnotation.accessType.name,
        roleTypes = emptyList()
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
    
    // Determine parameter type with more specific info
    val protoType = determineProtoType(field, parameter)
    
    // Get validation annotations
    val validations = extractValidations(field)
    
    // Extract authorization info if present
    val authorizationInfo = extractAuthorizationInfo(parameter)
    
    return ParameterMetadata(
        name = parameter.name,
        fieldName = field.name,
        javaType = field.type.name,
        commandType = parameter.type.name,
        protoType = protoType.type,
        protoTypeAnnotation = protoType.annotation,
        description = parameter.description,
        required = parameter.required,
        since = parameter.since,
        entityType = parameter.annotationClass.simpleName,
        validations = validations,
        authorization = authorizationInfo,
        shouldDisplay = true
    )
}

/**
 * Determine the appropriate protobuf type for a parameter
 */
fun determineProtoType(field: Field, parameter: Parameter): ProtoTypeInfo {
    val type = parameter.type
    
    // Basic type mapping
    val baseType = when (type) {
        CommandType.STRING -> "string"
        CommandType.BOOLEAN -> "bool"
        CommandType.INTEGER, CommandType.SHORT -> "int32"
        CommandType.LONG -> "int64"
        CommandType.FLOAT -> "float"
        CommandType.DOUBLE -> "double"
        CommandType.LIST -> "repeated string" // This is simplified
        CommandType.MAP -> "map<string, string>" // This is simplified
        CommandType.DATE -> "string" // Use string for dates
        CommandType.UUID -> "string" // Use string with annotation for UUIDs
        else -> "string"
    }
    
    // Check for special field types based on name or other hints
    val fieldName = parameter.name.lowercase()
    val annotation = when {
        // UUID type
        type == CommandType.UUID -> "CUSTOM_STRING_TYPE_UUID"
		parameter.collectionType == CommandType.STRING -> "okay"
//		type == CommandType.STRING ->


        
        // IP addresses
        fieldName.contains("ip") && !fieldName.contains("ipv6") -> "CUSTOM_STRING_TYPE_IP_V4_ADDRESS"
        fieldName.contains("ipv6") -> "CUSTOM_STRING_TYPE_IP_V6_ADDRESS"
        
        // CIDR blocks
        fieldName.contains("cidr") -> "CUSTOM_STRING_TYPE_CIDR_BLOCK"
        
        // MAC addresses
        fieldName.contains("mac") -> "CUSTOM_STRING_TYPE_MAC_ADDRESS"
        
        // Time/Date fields
        fieldName.contains("time") || fieldName.contains("date") -> "CUSTOM_STRING_TYPE_TIMESTAMP"
        
        // URLs
        fieldName.contains("url") -> "CUSTOM_STRING_TYPE_URL"
        
        // Emails
        fieldName.contains("email") -> "CUSTOM_STRING_TYPE_EMAIL"
        
        // Default - no annotation needed
        else -> null
    }
    
    return ProtoTypeInfo(baseType, annotation)
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