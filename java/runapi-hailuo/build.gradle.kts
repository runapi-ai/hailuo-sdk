plugins {
  `java-library`
  `maven-publish`
}

extra["runapiSlug"] = "hailuo"

description = "RunAPI Hailuo Java SDK for Hailuo workflows."

java {
  withSourcesJar()
  withJavadocJar()
}

dependencies {
  api("ai.runapi:runapi-core:0.4.1")

  testImplementation(platform("org.junit:junit-bom:5.10.3"))
  testImplementation("org.junit.jupiter:junit-jupiter")
}

publishing {
  publications {
    create<MavenPublication>("mavenJava") {
      from(components["java"])
      artifactId = "runapi-hailuo"
      pom {
        name = "RunAPI Hailuo Java SDK"
        description = "RunAPI Hailuo Java SDK for Hailuo workflows."
        url = "https://runapi.ai/models/hailuo"
        licenses {
          license {
            name = "Apache License, Version 2.0"
            url = "https://www.apache.org/licenses/LICENSE-2.0"
          }
        }
        developers {
          developer {
            id = "runapi"
            name = "RunAPI"
            email = "contact@runapi.ai"
          }
        }
        scm {
          url = "https://github.com/runapi-ai/hailuo-sdk"
          connection = "scm:git:https://github.com/runapi-ai/hailuo-sdk.git"
          developerConnection = "scm:git:ssh://git@github.com/runapi-ai/hailuo-sdk.git"
        }
      }
    }
  }
}
