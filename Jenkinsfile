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
                git branch: "DevMaster", credentialsId: "fc8e020b15949259b6ca27c56ad3e0d32ba815b852ac31f397d35aa148f04d81", url: "https://github.com/tanawitpat/go-my-resume.git"
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
                sh "docker container rm -f go-my-resume"
                sh """
                docker run -d \
                --network my-resume \
                --name go-my-resume \
                --restart always \
                -p 40000:8050 \
                go-my-resume
                """
            }
        }
    }
}
