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

	val outDirLoc = File(outdir)
	outDirLoc.deleteRecursively()
	outDirLoc.mkdirs()

	val (cmds, resps) = generateMetadataFromPackage("org.apache.cloudstack.api")

	val enumTypes = EnumExtractor.extractEnumTypes("org.apache.cloudstack", "com.cloud")


	val groupedCommands = cmds.associateBy { it.className }

	val groupedResponses = resps.associateBy { it.className }

	val enumgroups = enumTypes.associateBy { it.className }


	File(outDirLoc, "cloudstack-api-command-metadata.json").writeText(objectWriter.writeValueAsString(groupedCommands))


	File(outDirLoc, "cloudstack-api-response-metadata.json").writeText(objectWriter.writeValueAsString(groupedResponses))


	File(outDirLoc, "cloudstack-api-enum-metadata.json").writeText(objectWriter.writeValueAsString(enumgroups))

	println()


	println("\nCloudStack API Generator Completed ")

	println("Output directory: ${outDirLoc.absolutePath}")
	println("Generated files:")
	println("  - cloudstack-api-command-metadata.json (${groupedCommands.size} commands)")
	println("  - cloudstack-api-response-metadata.json (${groupedResponses.size} responses)")
	println("  - cloudstack-api-enum-metadata.json (${enumgroups.size} enums)")
}



fun buildPrettyJsonObjectWriter() : ObjectWriter {
	val objectMapper = ObjectMapper().registerKotlinModule()
	val indenter = DefaultIndenter("\t", DefaultIndenter.SYS_LF)
	val printer = DefaultPrettyPrinter()
	printer.indentObjectsWith(indenter)
	printer.indentArraysWith(indenter)
	return objectMapper.writer(printer)
}