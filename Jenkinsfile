pipeline {
    agent any

    environment {
        DOCKER_REGISTRY_URL = "ghcr.io"
    }

    stages {

        stage('SonarQube analysis') {
            def scannerHome = tool 'SonarScanner 4.0';
            withSonarQubeEnv() { // If you have configured more than one global server connection, you can specify its name
                sh "sonar-scanner \
                    -Dsonar.projectKey=local \
                    -Dsonar.sources=. \
                    -Dsonar.host.url=http://sonarqube.local.com \
                    -Dsonar.token=sqp_04943cb1829d2ed4ed9e8a4fa54d39cf88ca5a1c"
            }
        }

        stage('Login repo') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: 'docker-login-credentials', usernameVariable: 'USER_LOGIN', passwordVariable: 'TOKEN_LOGIN')]) {
                        sh "docker login ${DOCKER_REGISTRY_URL} -u ${USER_LOGIN} -p ${TOKEN_LOGIN}"
                    }
               }
            }
        }

        stage('Build, Push, and Deploy') {
            steps {
                sh 'docker-compose build && docker-compose push && docker-compose up -d --build'
            }
        }

        stage('Countdown') {
            steps {
                script {
                for (int i = 60; i >= 0; i--) {
                    echo "เวลาที่เหลือ: ${i} วินาที"
                    sleep 1
                }
                echo "หมดเวลาแล้ว!"
                }
            }
        }

        stage('Remove') {
            steps {
                script {
                sh "docker-compose down --rmi all"
                }
            }
        }
    }
}
