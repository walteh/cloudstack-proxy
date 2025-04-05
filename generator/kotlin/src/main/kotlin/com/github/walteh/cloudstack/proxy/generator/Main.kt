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


/**
 * Main entry point for the CloudStack API schema generator
 */
fun main(args: Array<String>) {
	println("CloudStack API Generator Starting...")

	// Initialize Jackson ObjectMapper with Kotlin module
	val objectMapper = ObjectMapper().registerKotlinModule()


	val configBuilder = SchemaGeneratorConfigBuilder(SchemaVersion.DRAFT_7)
		.with(
			Option.VALUES_FROM_CONSTANT_FIELDS,
//			Option.PUBLIC_STATIC_FIELDS,
			Option.PUBLIC_NONSTATIC_FIELDS,
//			Option.NONPUBLIC_STATIC_FIELDS,
			Option.NONPUBLIC_NONSTATIC_FIELDS_WITH_GETTERS,
			Option.NONPUBLIC_NONSTATIC_FIELDS_WITHOUT_GETTERS,
			Option.TRANSIENT_FIELDS,
//			Option.STATIC_METHODS,
//			Option.VOID_METHODS,
//			Option.GETTER_METHODS,
//			Option.NONSTATIC_NONVOID_NONGETTER_METHODS,
			Option.ADDITIONAL_FIXED_TYPES,
			Option.SIMPLIFIED_ENUMS,
			Option.SIMPLIFIED_OPTIONALS,
			Option.DEFINITIONS_FOR_ALL_OBJECTS,
			Option.NULLABLE_FIELDS_BY_DEFAULT,
			Option.NULLABLE_METHOD_RETURN_VALUES_BY_DEFAULT,
			Option.EXTRA_OPEN_API_FORMAT_VALUES,
			Option.FLATTENED_ENUMS,
			Option.FLATTENED_SUPPLIERS,
			Option.FLATTENED_OPTIONALS,
			Option.ALLOF_CLEANUP_AT_THE_END,
			Option.DUPLICATE_MEMBER_ATTRIBUTE_CLEANUP_AT_THE_END,
			Option.NULLABLE_ALWAYS_AS_ANYOF,
			Option.SCHEMA_VERSION_INDICATOR,
//			Option.PLAIN_DEFINITION_KEYS,
			Option.STRICT_TYPE_INFO
		)
		.without(
			Option.INLINE_ALL_SCHEMAS,
			Option.EXTRA_OPEN_API_FORMAT_VALUES,
			Option.NONPUBLIC_STATIC_FIELDS,
//			Option.NONPUBLIC_NONSTATIC_FIELDS_WITH_GETTERS,
//			Option.NONPUBLIC_NONSTATIC_FIELDS_WITHOUT_GETTERS,
			Option.STATIC_METHODS,
			Option.VOID_METHODS,
			Option.GETTER_METHODS,
			Option.TRANSIENT_FIELDS,
//			Option.NONSTATIC_NONVOID_NONGETTER_METHODS,
//			Option.NULLABLE_FIELDS_BY_DEFAULT,
			Option.NULLABLE_METHOD_RETURN_VALUES_BY_DEFAULT,
		)
	val config = configBuilder.with(ModuleImpl()).build()
	val generator = SchemaGenerator(config)

	val prettyPrinter = DefaultPrettyPrinter()

	val indenter = DefaultIndenter("\t", DefaultIndenter.SYS_LF)
	val printer = DefaultPrettyPrinter()
	printer.indentObjectsWith(indenter)
	printer.indentArraysWith(indenter)

	// format with tabs 
//	objectMapper.writerWithDefaultPrettyPrinter().writeValueAsString(objectMapper.readTree("{}"))
	// Find all classes in com.user package
	val userClasses = getClassesForPackage("org.apache.cloudstack.api")
	println("\nFound ${userClasses.size} classes in com.user package:")
	userClasses.forEach { println(it) }

	var schemas = mutableMapOf<String, JsonNode>()
	userClasses.forEach {
			try {
				schemas[it.typeName] = generator.generateSchema(it)
				println("successfully processed: ${it.typeName}")
			} catch (e: Throwable) {
				println("Error generating schema for ${it.typeName}: ${e.message}")

			}

	}
	// schemas.forEach { println(it) }


	val outputDir = File("./gen/cloudstack-api-json-schemas")
	outputDir.mkdirs()

	schemas.forEach { e ->
		println("${e.key}, ${e.value.get("title")}") }

	schemas.forEach {
//		if (it.id != null) {
			val jsonSchema = objectMapper.writer(printer).writeValueAsString(it.value)
//			val fileName = it.id.removePrefix("urn:jsonschema:").replace(":", ".")
			File(outputDir, "${it.key}.json").writeText(jsonSchema)
//		}
		// write to file
	}

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
fun getClassesForPackage(pckgname: String): ArrayList< java.lang.reflect.Type> {
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


/**
 * Example created in response to [#248](https://github.com/victools/jsonschema-generator/issues/248).
 * <br></br>
 * For use via the Maven plugin where schemas for all types in a given package will be generated and should reference each other.
 */
//class ExternalRefPackageExample : SchemaGenerationExampleInterface {
//	override fun generateSchema(): ObjectNode {
//		val config = SchemaGeneratorConfigBuilder(SchemaVersion.DRAFT_2020_12, OptionPreset.PLAIN_JSON)
//			.with(ModuleImpl())
//			.build()
//		return SchemaGenerator(config).generateSchema(Example::class.java)
//	}
//


class ModuleImpl : Module {
	override fun applyToConfigBuilder(builder: SchemaGeneratorConfigBuilder) {
		val definitionProvider = SchemaRefDefinitionProvider(
			PACKAGE_FOR_EXTERNAL_REFS,
			EXTERNAL_REF_PREFIX, EXTERNAL_REF_SUFFIX
		)
		builder.forTypesInGeneral()
			.withCustomDefinitionProvider(definitionProvider)
			.withIdResolver { scope: TypeScope ->
				if (definitionProvider.isMainType(scope.type))
					definitionProvider.getExternalRef(scope.type)
				else
					null
			}
			.withCustomDefinitionProvider(EnumMapDefinitionProvider(Function<Enum<*>, String> { enumValue: Enum<*> ->
				enumValue.name.lowercase(
					Locale.getDefault()
				)
			}))
	}

	internal class SchemaRefDefinitionProvider(
		private val packageForExternalRefs: String,
		private val externalRefPrefix: String,
		private val externalRefSuffix: String
	) :
		CustomDefinitionProviderV2 {
		private var mainType: Class<*>? = null

		override fun provideCustomSchemaDefinition(
			javaType: ResolvedType,
			context: SchemaGenerationContext
		): CustomDefinition? {
			val erasedType = javaType.erasedType
			if (this.mainType == null) {
				this.mainType = erasedType
			} else if (!this.isMainType(javaType) && erasedType.getPackage() != null && erasedType.getPackage().name.startsWith(
					this.packageForExternalRefs
				)
			) {
				val schema = context.generatorConfig.createObjectNode()
					.put(context.getKeyword(SchemaKeyword.TAG_REF), this.getExternalRef(javaType))
				return CustomDefinition(
					schema,
					CustomDefinition.INLINE_DEFINITION,
					CustomDefinition.INCLUDING_ATTRIBUTES
				)
			}
			return null
		}

		fun isMainType(javaType: ResolvedType): Boolean {
			return this.mainType == javaType.erasedType
		}

		fun getExternalRef(javaType: ResolvedType): String {
			return this.externalRefPrefix + javaType.erasedType.name + this.externalRefSuffix
		}

		override fun resetAfterSchemaGenerationFinished() {
			this.mainType = null
		}
	}

	companion object {
		private const val PACKAGE_FOR_EXTERNAL_REFS = "com.cloud"
		private const val EXTERNAL_REF_PREFIX = "file://"
		private const val EXTERNAL_REF_SUFFIX = ".json"
	}
}


/**
 * Example created in response to [#376](https://github.com/victools/jsonschema-generator/discussions/376).
 * <br></br>
 * Representing a Map containing Enum values as keys.
 */
//class EnumMapExample : SchemaGenerationExampleInterface {
//	override fun generateSchema(): ObjectNode {
//		val configBuilder = SchemaGeneratorConfigBuilder(SchemaVersion.DRAFT_2020_12, OptionPreset.PLAIN_JSON)
//		configBuilder.forTypesInGeneral()
//			.withCustomDefinitionProvider(EnumMapDefinitionProvider(Function<Enum<*>, String> { enumValue: Enum<*> ->
//				enumValue.name.lowercase(
//					Locale.getDefault()
//				)
//			}))
//		val config = configBuilder.build()
//		val generator = SchemaGenerator(config)
//		return generator.generateSchema(
//			MutableMap::class.java,
//			Job::class.java,
//			JobDetails::class.java
//		)
//	}
//
//
//
//	internal enum class Job {
//		BLACKSMITH, FARMER
//	}
//
//	internal class JobDetails {
//		var trainingCosts: Map<TrainingResource, Int>? = null
//	}
//
//	internal enum class TrainingResource {
//		WOOD, FOOD, STONE
//	}
//}

class EnumMapDefinitionProvider(private val propertyNameSupplier: Function<Enum<*>, String>) :
	CustomDefinitionProviderV2 {
	override fun provideCustomSchemaDefinition(
		targetType: ResolvedType,
		context: SchemaGenerationContext
	): CustomDefinition? {
		val keyType = context.typeContext.getTypeParameterFor(
			targetType,
			MutableMap::class.java, 0
		)
		if (keyType == null || !keyType.isInstanceOf(Enum::class.java)) {
			// only consider Maps with an Enum as key
			return null
		}
		val valueType = Optional.ofNullable(
			context.typeContext.getTypeParameterFor(
				targetType,
				MutableMap::class.java, 1
			)
		)
			.orElseGet {
				context.typeContext.resolve(
					Any::class.java
				)
			}
		val customSchema = context.generatorConfig.createObjectNode()
		val propertiesNode = context.generatorConfig.createObjectNode()
		customSchema.set<JsonNode>(context.getKeyword(SchemaKeyword.TAG_PROPERTIES), propertiesNode)
		Stream.of(*(keyType.erasedType as Class<out Enum<*>?>).enumConstants)
			.map<String>(this.propertyNameSupplier)
			.forEach { propertyName: String? ->
				propertiesNode.set<JsonNode>(
					propertyName,
					context.createStandardDefinitionReference(valueType, this)
				)
			}
		return CustomDefinition(customSchema)
	}
}