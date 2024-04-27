pipeline {
    agent any

    environment {
        DOCKER_REGISTRY_URL = "ghcr.io"
    }

    stages {
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