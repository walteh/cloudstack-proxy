package com.github.walteh.cloudstack.proxy.generator

import org.apache.cloudstack.api.APICommand
import java.io.File
import java.net.URLClassLoader
import java.util.jar.JarFile

/**
 * Generate command metadata JSON files from a JAR file
 */
fun generateMetadataFromJar(cloudstackJarPath: String): List<GroupMetadata?> {
    println("Starting command metadata extraction from: $cloudstackJarPath")
    
    // Load command classes from JAR
    val commandClasses = loadCommandClasses(cloudstackJarPath)
    
    println("Found ${commandClasses.size} command classes")
    
    // Extract metadata and write JSON files
   val cmds = extractCommandsToJson(commandClasses)

    println("Completed metadata extraction")

	return cmds
}

/**
 * Generate metadata using classes found in a package
 */
fun generateMetadataFromPackage(packageName: String): List<GroupMetadata?> {
    println("Starting command metadata extraction from package: $packageName")
    
    // Get classes from the specified package
    val classes = getClassesForPackage(packageName)
    
    // Filter for command classes (those with @APICommand annotation)
    val commandClasses = classes.filterIsInstance<Class<*>>()
        .filter { clazz -> 
            try {
                clazz.isAnnotationPresent(APICommand::class.java)
            } catch (e: Exception) {
                false
            }
        }
    
    println("Found ${commandClasses.size} command classes")
    
    // Extract metadata and write JSON files
   val cmds =  extractCommandsToJson(commandClasses)
    
    println("Completed metadata extraction")

	return cmds
}

/**
 * Load command classes from a CloudStack JAR
 */
fun loadCommandClasses(jarPath: String): List<Class<*>> {
    val jarFile = JarFile(jarPath)
    val entries = jarFile.entries()
    val urls = arrayOf(File(jarPath).toURI().toURL())
    val classLoader = URLClassLoader(urls, Thread.currentThread().contextClassLoader)
    
    val commandClasses = mutableListOf<Class<*>>()
    
    // Process JAR entries
    while (entries.hasMoreElements()) {
        val entry = entries.nextElement()
        if (entry.name.endsWith(".class")) {
            val className = entry.name.replace("/", ".").removeSuffix(".class")
            try {
                val clazz = classLoader.loadClass(className)
                
                // Check if this is a command class (has @APICommand annotation)
                if (clazz.isAnnotationPresent(APICommand::class.java)) {
                    commandClasses.add(clazz)
                }
            } catch (e: Exception) {
                // Skip problematic classes
                println("Skipping class $className: ${e.message}")
            }
        }
    }
    
    return commandClasses
}

///**
// * Main entry point
// */
//fun main(args: Array<String>) {
//    if (args.isEmpty()) {
//        println("Usage:")
//        println("  jar <cloudstack-jar-path> <output-dir>")
//        println("  package <package-name> <output-dir>")
//        return
//    }
//
//
////	args = listOf("package", )
//    var cmds = listOf<GroupMetadata?>()
//    when (args[0]) {
//        "jar" -> {
//            if (args.size < 3) {
//                println("Missing arguments for JAR mode")
//                println("Usage: jar <cloudstack-jar-path> <output-dir>")
//                return
//            }
//            val cloudstackJarPath = args[1]
//            val outputDir = File(args[2])
//            cmds = generateMetadataFromJar(cloudstackJarPath)
//        }
//        "package" -> {
//            if (args.size < 3) {
//                println("Missing arguments for package mode")
//                println("Usage: package <package-name> <output-dir>")
//                return
//            }
//            val packageName = args[1]
//            val outputDir = File(args[2])
//           	cmds = generateMetadataFromPackage(packageName)
//        }
//        else -> {
//            println("Unknown mode: ${args[0]}")
//            println("Usage:")
//            println("  jar <cloudstack-jar-path> <output-dir>")
//            println("  package <package-name> <output-dir>")
//        }
//
//
//    }
//}