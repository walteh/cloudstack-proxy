package com.github.walteh.cloudstack.proxy.generator

import org.reflections.Reflections
import org.reflections.scanners.SubTypesScanner
import org.reflections.scanners.TypeAnnotationsScanner
import org.reflections.util.ClasspathHelper
import org.reflections.util.ConfigurationBuilder
import org.reflections.util.FilterBuilder
import java.util.LinkedHashMap

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
        "com.cloud.network"
    )
    
    /**
     * Extract all enum types from the specified packages
     */
    fun extractEnumTypes(): List<EnumTypeMetadata> {
        val enumTypes = mutableListOf<EnumTypeMetadata>()
        
        // Configure Reflections to scan the target packages
        val reflections = buildReflections()
        
        // Get all enum classes in the packages
        val enumClasses = reflections.getSubTypesOf(Enum::class.java)
        
        // Extract metadata from each enum class
        for (enumClass in enumClasses) {
            try {
                val metadata = extractEnumMetadata(enumClass)
                enumTypes.add(metadata)
            } catch (e: Exception) {
                println("Error extracting metadata from enum class: ${enumClass.name} - ${e.message}")
            }
        }
        
        // Also look for nested enum classes that might not be directly caught
        val allClasses = reflections.getSubTypesOf(Any::class.java)
        for (clazz in allClasses) {
            try {
                // Check if this class has nested enum classes
                for (nestedClass in clazz.declaredClasses) {
                    if (nestedClass.isEnum) {
                        val metadata = extractEnumMetadata(nestedClass as Class<out Enum<*>>)
                        enumTypes.add(metadata)
                    }
                }
            } catch (e: Exception) {
                // Skip classes that cause errors
            }
        }
        
        return enumTypes.sortedBy { it.className }
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
            val extraInfo = extractExtraInfo(enumValue)
            
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
     * Extract extra information from an enum constant through reflection
     */
    private fun extractExtraInfo(enumValue: Enum<*>): Map<String, String> {
        val extraInfo = LinkedHashMap<String, String>()
        
        // Get all methods that might contain extra info
        val methods = enumValue.javaClass.methods
        for (method in methods) {
            // Skip inherited methods from Enum or Object
            if (method.declaringClass == Enum::class.java || method.declaringClass == Object::class.java) {
                continue
            }
            
            // Skip methods that require parameters
            if (method.parameterCount > 0) {
                continue
            }
            
            try {
                // Try to get method value
                val value = method.invoke(enumValue)
                if (value != null) {
                    extraInfo[method.name] = value.toString()
                }
            } catch (e: Exception) {
                // Skip methods that throw exceptions
            }
        }
        
        return extraInfo
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