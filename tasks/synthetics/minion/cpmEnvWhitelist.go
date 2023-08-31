package minion

// Expected ENV variables should be uppercase
// Leave out expected ENV variables that have sensitive values:
// "MINION_VSE_PASSPHRASE"
// "MINION_API_PROXY_AUTH"
// "MINION_PRIVATE_LOCATION_KEY"
// "MINION_USER_DEFINED_ENV_VARIABLES"
var CPMenvWhitelist = []string{
	//Variables default to the CPM docker image
	"PATH",
	"JAVA_VERSION_MAJOR",
	"JAVA_VERSION_MINOR",
	"JAVA_VERSION_BUILD",
	"JAVA_PACKAGE",
	"JAVA_JCE",
	"JAVA_HOME",
	"GLIBC_REPO",
	"GLIBC_VERSION",
	"LANG",
	"WORKDIR",
	"MINION_JAR",
	"DEFAULT_LOCALE_CFG_FILE",
	"MINION_PROVIDER",
	"MINION_USER",
	"MINION_GROUP",
	//User Configurable variables below
	"MINION_LOG_LEVEL",
	"DOCKER_API_VERSION",
	"DOCKER_HOST",
	"MINION_API_ENDPOINT",
	"MINION_DOCKER_RUNNER_REGISTRY_ENDPOINT",
	"MINION_API_PROXY",
	"MINION_API_PROXY_SELF_SIGNED_CERT",
	"MINION_CHECK_TIMEOUT",
	"MINION_DOCKER_API_VERSION",
	"MINION_DOCKER_HOST",
	"MINION_DOCKER_RUNNER_APPARMOR",
	"MINION_JVM_MB",
	"MINION_JVM_OPTS",
}
