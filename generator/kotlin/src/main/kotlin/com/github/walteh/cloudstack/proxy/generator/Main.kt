package com.github.walteh.cloudstack.proxy.generator

// CloudStack API imports
//import kotlin.reflect.full.declaredClasses
import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.module.jsonSchema.JsonSchemaGenerator
import com.fasterxml.jackson.module.kotlin.registerKotlinModule
import java.io.File
import java.io.IOException
import java.io.UnsupportedEncodingException
import java.lang.reflect.Method
import java.net.*
import java.util.*
import java.util.jar.JarEntry
import java.util.jar.JarFile


/**
 * Main entry point for the CloudStack API schema generator
 */
fun main(args: Array<String>) {
	println("CloudStack API Generator Starting...")

	// Initialize Jackson ObjectMapper with Kotlin module
	val objectMapper = ObjectMapper().registerKotlinModule()
	val schemaGenerator = JsonSchemaGenerator(objectMapper)

	// Find all classes in com.user package
	val userClasses = getClassesForPackage("com.cloud")
	println("\nFound ${userClasses.size} classes in com.user package:")
	userClasses.forEach { println(it) }

	println("\nCloudStack API Generator Completed")
}


/**
 * Private helper method
 *
 * @param directory
 * The directory to start with
 * @param pckgname
 * The package name to search for. Will be needed for getting the
 * Class object.
 * @param classes
 * if a file isn't loaded but still is in the directory
 * @throws ClassNotFoundException
 */
@Throws(ClassNotFoundException::class)
private fun checkDirectory(
	directory: File, pckgname: String,
	classes: ArrayList<Class<*>>
) {
	var tmpDirectory: File

	if (directory.exists() && directory.isDirectory) {
		val files = directory.list()

		for (file in files!!) {
			if (file.endsWith(".class")) {
				try {
					classes.add(
						Class.forName(
							(pckgname + '.'
								+ file.substring(0, file.length - 6))
						)
					)
				} catch (e: NoClassDefFoundError) {
					// do nothing. this class hasn't been found by the
					// loader, and we don't care.
				}
			} else if ((File(directory, file).also { tmpDirectory = it })
					.isDirectory
			) {
				checkDirectory(tmpDirectory, "$pckgname.$file", classes)
			}
		}
	}
}

/**
 * Private helper method.
 *
 * @param connection
 * the connection to the jar
 * @param pckgname
 * the package name to search for
 * @param classes
 * the current ArrayList of all classes. This method will simply
 * add new classes.
 * @throws ClassNotFoundException
 * if a file isn't loaded but still is in the jar file
 * @throws IOException
 * if it can't correctly read from the jar file.
 */
@Throws(ClassNotFoundException::class, IOException::class)
private fun checkJarFile(
	connection: java.net.JarURLConnection,
	pckgname: String,
	classes: ArrayList<Class<*>>
) {
	val jarFile: JarFile = connection.getJarFile()
	println("checking jar file: ${jarFile.toString()}")
	val entries: Enumeration<JarEntry> = jarFile.entries()
	var name: String

	var jarEntry: JarEntry? = null
	while (entries.hasMoreElements()
		&& ((entries.nextElement().also { jarEntry = it }) != null)
	) {

		name = jarEntry?.name ?: ""

		if (name.contains(".class")) {
			name = name.substring(0, name.length - 6).replace('/', '.')

			if (name.contains(pckgname)) {
				println("$name, $pckgname")
				classes.add(Class.forName(name))
			}
		}
	}
}

/**
 * Attempts to list all the classes in the specified package as determined
 * by the context class loader
 *
 * @param pckgname
 * the package name to search
 * @return a list of classes that exist within that package
 * @throws ClassNotFoundException
 * if something went wrong
 */
@Throws(ClassNotFoundException::class)
fun getClassesForPackage(pckgname: String): ArrayList<Class<*>> {
	val classes = ArrayList<Class<*>>()

	try {
		val cld = Thread.currentThread()
			.contextClassLoader

		if (cld == null) throw ClassNotFoundException("Can't get class loader.")

		val resources = cld.getResources(
			pckgname
				.replace('.', '/')
		)
		var connection: URLConnection

		var url: URL? = null
		while (resources.hasMoreElements()
			&& ((resources.nextElement().also { url = it }) != null)
		) {
			try {
				connection = url!!.openConnection()

				if (connection is java.net.JarURLConnection) {
					checkJarFile(
						connection as java.net.JarURLConnection, pckgname,
						classes
					)
				} else if (url!!.getProtocol().equals("file")) {
					try {
						checkDirectory(
							File(
								URLDecoder.decode(
									url!!.path,
									"UTF-8"
								)
							), pckgname, classes
						)
					} catch (ex: UnsupportedEncodingException) {
						throw ClassNotFoundException(
							pckgname
								+ " does not appear to be a valid package (Unsupported encoding)",
							ex
						)
					}
				} else throw ClassNotFoundException(
					(pckgname + " ("
						+ url!!.path
						+ ") does not appear to be a valid package")
				)
			} catch (ioex: IOException) {
				throw ClassNotFoundException(
					"IOException was thrown when trying to get all resources for "
						+ pckgname, ioex
				)
			}
		}
	} catch (ex: NullPointerException) {
		throw ClassNotFoundException(
			pckgname
				+ " does not appear to be a valid package (Null pointer exception)",
			ex
		)
	} catch (ioex: IOException) {
		throw ClassNotFoundException(
			"IOException was thrown when trying to get all resources for "
				+ pckgname, ioex
		)
	}

	return classes
}