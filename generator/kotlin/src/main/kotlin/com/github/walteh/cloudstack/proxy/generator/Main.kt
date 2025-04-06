package com.github.walteh.cloudstack.proxy.generator

// CloudStack API imports
//import kotlin.reflect.full.declaredClasses
import com.fasterxml.classmate.ResolvedType
import com.fasterxml.jackson.core.util.DefaultIndenter
import com.fasterxml.jackson.core.util.DefaultPrettyPrinter
import com.fasterxml.jackson.databind.JsonNode
import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.databind.ObjectWriter
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


/**
 * Main entry point for the CloudStack API schema generator
 */
suspend fun main(args: Array<String>) {
	println("CloudStack API Generator Starting...")



	val objectWriter = buildPrettyJsonObjectWriter()

	val userClasses = getClassesForPackage("org.apache.cloudstack.api")

//	val schemas = convertToJsonSchemas(userClasses) { !it.endsWith("Cmd") }




	// val schemaOutputDir = File("./gen/cloudstack-api-json-schemas")
	// schemaOutputDir.mkdirs()

	// schemas.forEach {
	// 	if (it.error == null) {
	// 		val jsonSchema = objectWriter.writeValueAsString(it.schema)
	// 		File(schemaOutputDir, "${it.name}.json").writeText(jsonSchema)
	// 	}
	// }

	val cmds = generateMetadataFromPackage("org.apache.cloudstack.api")

	val metadataOutputDir = File("./gen/cloudstack-api-metadata")
	metadataOutputDir.deleteRecursively()
	metadataOutputDir.mkdirs()

	cmds.forEach { 
		if (it != null) {
			val json = objectWriter.writeValueAsString(it)
			File(metadataOutputDir, "${it.name}.json").writeText(json)
		}
	}


	println("\nCloudStack API Generator Completed")
}



fun buildPrettyJsonObjectWriter() : ObjectWriter {
	val objectMapper = ObjectMapper().registerKotlinModule()
	val indenter = DefaultIndenter("\t", DefaultIndenter.SYS_LF)
	val printer = DefaultPrettyPrinter()
	printer.indentObjectsWith(indenter)
	printer.indentArraysWith(indenter)
	return objectMapper.writer(printer)
}