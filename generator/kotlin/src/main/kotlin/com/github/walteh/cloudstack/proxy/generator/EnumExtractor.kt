package com.github.walteh.cloudstack.proxy.generator

import java.util.LinkedHashMap
import java.lang.reflect.Method
import java.lang.reflect.Modifier
import java.lang.reflect.Type

/**
 * Data class to represent an Enum Type
 */
data class EnumTypeMetadata(
    val className: String,
    val simpleName: String,
    val packageName: String,
    val isNested: Boolean,
    val parentClass: String?,
    val values: List<EnumValueMetadata>
)

/**
 * Data class to represent an Enum Value
 */
data class EnumValueMetadata(
    val name: String,
    val ordinal: Int,
    val extras: Map<String, String> = emptyMap()
)

/**
 * Extracts all enum types from CloudStack packages
 */
object EnumExtractor {
    
    // Base package to start scanning from
    
    // Methods to skip in extras extraction
    private val methodsToSkip = setOf(
        "values", "valueOf", "getDeclaringClass", "clone", 
        "getClass", "hashCode", "toString", "notify", "notifyAll", "wait"
    )
    
    /**
     * Extract all enum types from CloudStack API
     */
    fun extractEnumTypes(vararg packageNames: String): List<EnumTypeMetadata> {
        val enumTypes = mutableMapOf<String, EnumTypeMetadata>()
        println("Starting enum extraction from CloudStack API...")
        
        // First, let's get all CloudStack classes starting from the base package
        val allClasses = packageNames.flatMap { packageName ->
            getClassesForPackage(packageName) { true }
                .filterIsInstance<Class<*>>()
        }
        

        
        // Combine all classes
        
        println("Found ${allClasses.size} total classes to check for enums")
        
        // Process each class to find enums
        processClassesForEnums(allClasses, enumTypes)
        
        println("Total unique enum types found: ${enumTypes.size}")
        return enumTypes.values.toList().sortedBy { it.className }
    }
    
    /**
     * Process all classes to find enums and nested enums
     */
    private fun processClassesForEnums(
        classes: List<Class<*>>, 
        enumTypes: MutableMap<String, EnumTypeMetadata>
    ) {
        // First, find all direct enum classes
        val directEnumClasses = classes.filter { it.isEnum }
        println("Found ${directEnumClasses.size} direct enum classes")
        
        // Process direct enum classes
        directEnumClasses.forEach { enumClass ->
            try {
                val metadata = extractEnumMetadata(enumClass as Class<out Enum<*>>)
                enumTypes[metadata.className] = metadata
            } catch (e: Exception) {
                println("Error extracting metadata from enum class: ${enumClass.name} - ${e.message}")
            }
        }
        
        // Then find all classes that might contain nested enums
        classes.forEach { clazz ->
            try {
                // Check declared classes for nested enums
                clazz.declaredClasses.filter { it.isEnum }.forEach { nestedEnum ->
                    try {
                        val metadata = extractEnumMetadata(nestedEnum as Class<out Enum<*>>)
                        enumTypes[metadata.className] = metadata
                    } catch (e: Exception) {
                        println("Error extracting metadata from nested enum: ${nestedEnum.name} - ${e.message}")
                    }
                }
                
                // Special handling for interface enums
                if (clazz.isInterface) {
                    clazz.declaredClasses.filter { it.isEnum }.forEach { nestedEnum ->
                        try {
                            val metadata = extractEnumMetadata(nestedEnum as Class<out Enum<*>>)
                            enumTypes[metadata.className] = metadata
                        } catch (e: Exception) {
                            println("Error extracting metadata from interface enum: ${nestedEnum.name} - ${e.message}")
                        }
                    }
                }
                
                // Check for enum fields which might be standalone enum types
                clazz.declaredFields.filter { 
                    Modifier.isStatic(it.modifiers) && it.type.isEnum 
                }.forEach { field ->
                    try {
                        val enumClass = field.type as Class<out Enum<*>>
                        if (!enumTypes.containsKey(enumClass.name)) {
                            val metadata = extractEnumMetadata(enumClass)
                            enumTypes[metadata.className] = metadata
                        }
                    } catch (e: Exception) {
                        // Skip fields that cause errors
                    }
                }
            } catch (e: Exception) {
                // Skip classes that cause errors
            }
        }
    }
    
    /**
     * Get all classes from a package with a filter
     */
    private fun getClassesForPackage(packageName: String, filter: (Type) -> Boolean): List<Type> {
        // Use the existing method from Custom.kt
        try {
            val allTypes = com.github.walteh.cloudstack.proxy.generator.getClassesForPackage(packageName)
            return allTypes.filter(filter)
        } catch (e: Exception) {
            println("Error loading classes from package $packageName: ${e.message}")
            return emptyList()
        }
    }
    
    /**
     * Group enums by package
     */
    fun groupEnumsByPackage(enums: List<EnumTypeMetadata>): Map<String, List<EnumTypeMetadata>> {
        return enums.groupBy { it.packageName }
    }
    
    /**
     * Extract metadata from an enum class
     */
    private fun extractEnumMetadata(enumClass: Class<out Enum<*>>): EnumTypeMetadata {
        val packageName = enumClass.packageName
        val isNested = enumClass.enclosingClass != null
        val parentClass = if (isNested) enumClass.enclosingClass.name else null
        
        // Get all enum constants
        val enumValues = enumClass.enumConstants ?: emptyArray()
        val valueMetadata = enumValues.map { enumValue ->
            // Extract extra info from the enum constant if available
            val extraInfo = extractUsefulExtraInfo(enumValue, enumClass)

			// if there is .getPossibleEvents on the extra stuff, sort it
			if (extraInfo.containsKey("getPossibleEvents")) {
				val extraEvents = extraInfo["getPossibleEvents"]?.removePrefix("[")?.removeSuffix("]")?.split(",")?.map { it.trim() }?.sorted()
				extraInfo["getPossibleEvents"] = "[${extraEvents?.joinToString(",")}]"
			}

			// create a new map with the extra info sorted by key
			val sortedExtraInfo = extraInfo.entries.sortedBy { it.key }

			// convert the list to a map
			val sortedExtraInfoMap = sortedExtraInfo.map { it.key to it.value }.toMap()
            
            EnumValueMetadata(
                name = enumValue.name,
                ordinal = enumValue.ordinal,
                extras = sortedExtraInfoMap
            )
        }
        
        return EnumTypeMetadata(
            className = enumClass.name,
            simpleName = enumClass.simpleName,
            packageName = packageName,
            isNested = isNested,
            parentClass = parentClass,
            values = valueMetadata
        )
    }
    
    /**
     * Extract useful information from an enum constant through reflection
     * Filters out common methods and focuses on valuable extras
     */
    private fun extractUsefulExtraInfo(enumValue: Enum<*>, enumClass: Class<out Enum<*>>): MutableMap<String, String> {
        val extraInfo = mutableMapOf<String, String>()
        
        // Get all declared methods in the enum class
        val declaredMethods = enumClass.declaredMethods
        
        // Get all methods from the enum value
        val methods = enumValue.javaClass.methods
        
        // First check declared methods which might include non-inherited methods
        for (method in declaredMethods) {
            extractInfoFromMethod(method, enumValue, extraInfo)
        }
        
        // Then check instance methods
        for (method in methods) {
            // Skip methods we've already processed and inherited methods from Enum/Object
            if (method.declaringClass == Enum::class.java || 
                method.declaringClass == Object::class.java || 
                methodsToSkip.contains(method.name)) {
                continue
            }
            
            extractInfoFromMethod(method, enumValue, extraInfo)
        }
        
        return extraInfo
    }
    
    private fun extractInfoFromMethod(method: Method, enumValue: Enum<*>, extraInfo: MutableMap<String, String>) {
        try {
            // Skip if method is in our skip list
            if (methodsToSkip.contains(method.name)) {
                return
            }
            
            // Skip methods that require parameters
            if (method.parameterCount > 0) {
                return
            }
            
            // Make private methods accessible
            if (!method.isAccessible) {
                method.isAccessible = true
            }
            
            // Try to get method value
            val value = method.invoke(enumValue)
            if (value != null) {
                val valueStr = value.toString()
                // Skip values that are clearly object references or not useful 
                if (!valueStr.contains("@") && 
                    !valueStr.startsWith("[L") && 
                    !valueStr.equals("0") && 
                    !valueStr.equals("1") && 
                    !valueStr.equals("2") && 
                    !valueStr.equals("3") && 
                    !valueStr.equals("4") && 
                    !valueStr.equals(enumValue.name)) {
                    extraInfo[method.name] = valueStr
                }
            }
        } catch (e: Exception) {
            // Skip methods that throw exceptions
        }
    }
    
    /**
     * Main method to run the extraction
     */
    // @JvmStatic
    // fun main(args: Array<String>) {
    //     val enumTypes = extractEnumTypes()
    //     println("Extracted ${enumTypes.size} enum types")
        
    //     // Group enums by package
    //     val enumsByPackage = groupEnumsByPackage(enumTypes)
    //     println("\nEnums by package:")
    //     enumsByPackage.forEach { (packageName, enums) ->
    //         println("  $packageName: ${enums.size} enums")
    //     }
        
    //     // Print all enum types for debugging
    //     enumTypes.forEach { enumType ->
    //         println("Enum: ${enumType.className}")
    //         println("  Values:")
    //         enumType.values.forEach { value ->
    //             println("    ${value.name} (${value.ordinal})")
    //             if (value.extras.isNotEmpty()) {
    //                 println("      Extra Info:")
    //                 value.extras.forEach { (key, extraValue) ->
    //                     println("        $key: $extraValue")
    //                 }
    //             }
    //         }
    //         println()
    //     }
    // }
} 