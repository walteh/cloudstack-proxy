package com.github.walteh.cloudstack.proxy.generator

// CloudStack API imports
//import kotlin.reflect.full.declaredClasses
import com.fasterxml.classmate.ResolvedType
import com.fasterxml.jackson.core.util.DefaultIndenter
import com.fasterxml.jackson.core.util.DefaultPrettyPrinter
import com.fasterxml.jackson.databind.JsonNode
import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.databind.node.JsonNodeFactory
import com.fasterxml.jackson.databind.node.ObjectNode
import com.fasterxml.jackson.module.kotlin.registerKotlinModule
import com.github.victools.jsonschema.generator.*
import kotlinx.coroutines.*
import kotlinx.coroutines.sync.Mutex
import kotlinx.coroutines.sync.withLock
import java.io.File
import java.io.IOException
import java.io.UnsupportedEncodingException
import java.net.*
import java.util.*
import java.util.concurrent.locks.Lock
import java.util.function.Function
import java.util.jar.JarEntry
import java.util.jar.JarFile
import java.util.stream.Stream
import kotlin.coroutines.CoroutineContext
import kotlinx.coroutines.*

//import kotlin.coroutines.




suspend fun convertToJsonSchemas(userClasses : ArrayList< java.lang.reflect.Type>, shouldGenerate: (String) -> Boolean = {true}): List<Schema> {

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



	// format with tabs 
//	objectMapper.writerWithDefaultPrettyPrinter().writeValueAsString(objectMapper.readTree("{}"))
	// Find all classes in com.user package
//	val userClasses = getClassesForPackage("org.apache.cloudstack.api")
	println("\nFound ${userClasses.size} classes in com.user package:")
	userClasses.forEach { println(it) }

	val dummynode = ObjectNode(JsonNodeFactory(true))
	val scope = coroutineScope {
		val jobs = userClasses.map {
			async {
				if (shouldGenerate(it.typeName)) {
					try {
						val s = generator.generateSchema(it)
						println("Successfully processed schema: ${it.typeName}")
						Schema(it.typeName, s, null)
					} catch (e: Throwable) {
						println("Error generating schema for ${it.typeName}: ${e.message}")
						Schema(it.typeName, dummynode, e)
					}
				} else {
					Schema(it.typeName, dummynode, Throwable("skipped"))
				}
			}
			}

		jobs
	}
//	var coroutines = mutableListOf<Job>()
//	var context = Co
	// process with coroutines


	// wait for all coroutines to finish`



	return scope.awaitAll()
}

data class Schema(val name: String, val schema: JsonNode, val error: Throwable?) {
	}

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
			.map {
				it?.name
			}
			.forEach { propertyName: String? ->
				propertiesNode.set<JsonNode>(
					propertyName,
					context.createStandardDefinitionReference(valueType, this)
				)
			}
		return CustomDefinition(customSchema)
	}
}