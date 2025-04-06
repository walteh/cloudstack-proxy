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
import kotlinx.cli.ArgParser
import kotlinx.cli.ArgType
import kotlinx.cli.required
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
	val parser = ArgParser	("cloudstack-api-generator")

	println("CloudStack API Generator Starting...")

    val outdir by parser.option(ArgType.String, shortName = "o", fullName = "out-dir", description = "Output directory").required()

	parser.parse(args)

//	val userClasses = getClassesForPackage("org.apache.cloudstack.api")
//
//	println("found ${userClasses.count()}")

//	val schemas = convertToJsonSchemas(userClasses) { !it.endsWith("Cmd") }


	val objectWriter = buildPrettyJsonObjectWriter()


	// val schemaOutputDir = File("./gen/cloudstack-api-json-schemas")
	// schemaOutputDir.mkdirs()

	// schemas.forEach {
	// 	if (it.error == null) {
	// 		val jsonSchema = objectWriter.writeValueAsString(it.schema)
	// 		File(schemaOutputDir, "${it.name}.json").writeText(jsonSchema)
	// 	}
	// }

	val cmds = generateMetadataFromPackage("org.apache.cloudstack.api")

	val outDirLoc = File(outdir)
	outDirLoc.deleteRecursively()
	outDirLoc.mkdirs()

	File(outDirLoc, "cloudstack-api-command-metadata.json").writeText(objectWriter.writeValueAsString(cmds))


	// Extract enum types
	println("\nExtracting enum types...")
	val enumTypes = EnumExtractor.extractEnumTypes()
	println("Found ${enumTypes.size} enum types")


	File(outDirLoc, "cloudstack-api-enum-metadata.json").writeText(objectWriter.writeValueAsString(enumTypes))


//	// Group enums by package
//	val enumsByPackage = enumTypes.groupBy { it.packageName }
//
//	enumsByPackage.forEach { (packageName, enums) ->
//		// Create a package-level JSON file
//		val packageShortName = packageName.split(".").last()
//		val json = objectWriter.writeValueAsString(enums)
//		File(enumOutputDir, "$packageShortName.json").writeText(json)
//		println("Wrote ${enums.size} enums to $packageShortName.json")
//	}
//
//	// Also write a consolidated file with all enums
//	val allEnumsJson = objectWriter.writeValueAsString(enumTypes)
//	File(enumOutputDir, "all_enums.json").writeText(allEnumsJson)
//	println("Wrote all ${enumTypes.size} enums to all_enums.json")

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