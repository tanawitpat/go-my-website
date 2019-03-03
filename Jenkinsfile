pipeline {
    agent {
        label "master"
    }
    options {
        skipDefaultCheckout true
    }
    stages {
        stage("PULL CODE") {
            steps {
                git branch: "DevMaster", credentialsId: "a88c4be69d5e46baf4f6c545e99306901a75ee5c4928b36063450db2fa55b910", url: "https://github.com/OuiPTN/go-my-resume.git"
                script{
                    gitCommit = sh (script: "git rev-parse HEAD | cut -c1-7", returnStdout: true).trim()
                    currentBuild.displayName = "#${env.BUILD_NUMBER}_${gitCommit}"
                }
            }
        }
        stage("BUILD IMAGE") {
            steps {
                sh "docker build . -t go-my-resume"
            }
        }
        stage("DEPLOY") {
            steps {
                sh """
                docker run --rm -d --network my-resume --name go-my-resume \
                -p 40000:8050 \
                go-my-resume
                """
            }
        }
    }
}
