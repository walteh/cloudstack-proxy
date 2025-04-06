package com.github.walteh.cloudstack.proxy.generator

import org.reflections.Reflections
import org.reflections.scanners.SubTypesScanner
import org.reflections.scanners.TypeAnnotationsScanner
import org.reflections.util.ClasspathHelper
import org.reflections.util.ConfigurationBuilder
import org.reflections.util.FilterBuilder
import java.util.LinkedHashMap
import java.lang.reflect.Method

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
 * Extracts all enum types from the org.apache.cloudstack.api package
 */
object EnumExtractor {
    
    private val basePackages = listOf(
        "org.apache.cloudstack.api",
        "org.apache.cloudstack.acl",
        "com.cloud.network",
        "com.cloud.dc",
        "com.cloud.user",
        "com.cloud.vm"
    )
    
    // Methods to skip in extras extraction
    private val methodsToSkip = setOf(
        "values", "valueOf", "getDeclaringClass", "clone", 
        "getClass", "hashCode", "toString", "notify", "notifyAll", "wait"
    )
    
    /**
     * Extract all enum types from the specified packages
     */
    fun extractEnumTypes(): List<EnumTypeMetadata> {
        val enumTypes = mutableMapOf<String, EnumTypeMetadata>()
        
        // Configure Reflections to scan the target packages
        val reflections = buildReflections()
        
        // Get all enum classes in the packages
        val enumClasses = reflections.getSubTypesOf(Enum::class.java)

        
        // Extract metadata from each enum class
        for (enumClass in enumClasses) {
            try {
                val metadata = extractEnumMetadata(enumClass)
                enumTypes[metadata.className] = metadata
            } catch (e: Exception) {
                println("Error extracting metadata from enum class: ${enumClass.name} - ${e.message}")
            }
        }


        
        // Also look for nested enum classes that might not be directly caught
        val allClasses = reflections.getSubTypesOf(Any::class.java)

		// add interface types to the enum types
		// allClasses += 
        for (clazz in allClasses) {
            try {
                // Check if this class has nested enum classes
                for (nestedClass in clazz.declaredClasses) {
                    if (nestedClass.isEnum) {
                        val metadata = extractEnumMetadata(nestedClass as Class<out Enum<*>>)
                        enumTypes[metadata.className] = metadata
                    }
                }
            } catch (e: Exception) {
                // Skip classes that cause errors
            }
        }
        
        return enumTypes.values.toList().sortedBy { it.className }
    }
    
    /**
     * Group enums by package
     */
    fun groupEnumsByPackage(enums: List<EnumTypeMetadata>): Map<String, List<EnumTypeMetadata>> {
        return enums.groupBy { it.packageName }
    }
    
    /**
     * Build the Reflections object configured to scan the target packages
     */
    private fun buildReflections(): Reflections {
        val classLoadersList = mutableListOf<ClassLoader>()
        classLoadersList.add(ClasspathHelper.contextClassLoader())
        classLoadersList.add(ClasspathHelper.staticClassLoader())
        
        val builder = ConfigurationBuilder()
        
        // Add URLs for all base packages
        basePackages.forEach { pkg ->
            builder.addUrls(ClasspathHelper.forPackage(pkg))
        }
        
        // Set class loaders and scanners
        builder.setClassLoaders(classLoadersList.toTypedArray())
            .addScanners(SubTypesScanner(false), TypeAnnotationsScanner())
        
        // Add filter to include only classes from the specified packages
        val filter = FilterBuilder()
        basePackages.forEach { pkg -> filter.includePackage(pkg) }
        builder.filterInputsBy(filter)
        
        return Reflections(builder)
    }
    
    /**
     * Extract metadata from an enum class
     */
    private fun extractEnumMetadata(enumClass: Class<out Enum<*>>): EnumTypeMetadata {
        val packageName = enumClass.packageName
        val isNested = enumClass.enclosingClass != null
        val parentClass = if (isNested) enumClass.enclosingClass.name else null
        
        // Get all enum constants
        val enumValues = enumClass.enumConstants
        val valueMetadata = enumValues.map { enumValue ->
            // Extract extra info from the enum constant if available
            val extraInfo = extractUsefulExtraInfo(enumValue, enumClass)
            
            EnumValueMetadata(
                name = enumValue.name,
                ordinal = enumValue.ordinal,
                extras = extraInfo
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
    private fun extractUsefulExtraInfo(enumValue: Enum<*>, enumClass: Class<out Enum<*>>): Map<String, String> {
        val extraInfo = LinkedHashMap<String, String>()
        
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
     * Write the extracted enum types to a file
     */
    fun writeToFile(filePath: String, enumTypes: List<EnumTypeMetadata>) {
        // Implementation depends on your file writing requirements
        // Could use a JSON serializer or other format
    }
    
    /**
     * Main method to run the extraction
     */
    @JvmStatic
    fun main(args: Array<String>) {
        val enumTypes = extractEnumTypes()
        println("Extracted ${enumTypes.size} enum types")
        
        // Group enums by package
        val enumsByPackage = groupEnumsByPackage(enumTypes)
        println("\nEnums by package:")
        enumsByPackage.forEach { (packageName, enums) ->
            println("  $packageName: ${enums.size} enums")
        }
        
        // Print all enum types for debugging
        enumTypes.forEach { enumType ->
            println("Enum: ${enumType.className}")
            println("  Values:")
            enumType.values.forEach { value ->
                println("    ${value.name} (${value.ordinal})")
                if (value.extras.isNotEmpty()) {
                    println("      Extra Info:")
                    value.extras.forEach { (key, extraValue) ->
                        println("        $key: $extraValue")
                    }
                }
            }
            println()
        }
    }
} 