package com.github.walteh.cloudstack.proxy.generator

// CloudStack API imports
//import kotlin.reflect.full.declaredClasses
import com.fasterxml.classmate.ResolvedType
import com.fasterxml.jackson.core.util.DefaultIndenter
import com.fasterxml.jackson.core.util.DefaultPrettyPrinter
import com.fasterxml.jackson.databind.JsonNode
import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.databind.node.ObjectNode
import com.fasterxml.jackson.module.kotlin.registerKotlinModule
import com.github.victools.jsonschema.generator.*
import java.io.File
import java.io.IOException
import java.io.UnsupportedEncodingException
import java.net.*
import java.util.*
import java.util.function.Function
import java.util.jar.JarEntry
import java.util.jar.JarFile
import java.util.stream.Stream





@Throws(ClassNotFoundException::class)
public fun getClassesForPackage(pckgname: String): ArrayList< java.lang.reflect.Type> {
	val classes = ArrayList< java.lang.reflect.Type>()

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






@Throws(ClassNotFoundException::class)
private fun checkDirectory(
	directory: File, pckgname: String,
	classes: ArrayList<java.lang.reflect.Type>
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


@Throws(ClassNotFoundException::class, IOException::class)
private fun checkJarFile(
	connection: java.net.JarURLConnection,
	pckgname: String,
	classes: ArrayList<java.lang.reflect.Type>
) {
	val jarFile: JarFile = connection.getJarFile()
	println("checking jar file: ${jarFile.toString()}")
	val entries: Enumeration<JarEntry> = jarFile.entries()
	var name: String

	var jarEntry: JarEntry? = null
	while (entries.hasMoreElements()
		&& ((entries.nextElement().also { jarEntry = it }) != null)
	) {

		name = (jarEntry?.name ?: "").replace('/', '.')

		if (name.contains(pckgname)) {

			println("$name, $pckgname")

			if (name.contains(".class")) {
				name = name.substring(0, name.length - 6)
				classes.add(Class.forName(name))

			} else if (name.contains(".enum")) {
				name = name.substring(0, name.length - 5)
				classes.add(Class.forName(name))

			}
		}
	}
}
