pipeline {
    agent any
    
    environment {
        // DockerHub credentials stored in Jenkins
        DOCKERHUB_CREDENTIALS = credentials('dockerhub-credentials')
        
        // DockerHub username
        DOCKERHUB_USERNAME = 'omarwael01'
        
        // Image names
        DATABASE_IMAGE = "${DOCKERHUB_USERNAME}/clinic-database"
        BACKEND_IMAGE = "${DOCKERHUB_USERNAME}/clinic-backend"
        FRONTEND_IMAGE = "${DOCKERHUB_USERNAME}/clinic-frontend"
        
        // Build number tag
        IMAGE_TAG = "${BUILD_NUMBER}"
        
        // API URL for frontend
        API_URL = 'http://localhost:8080'
    }
    
    stages {
        stage('Checkout') {
            steps {
                echo 'Checking out code from repository...'
                checkout scm
            }
        }
        
        stage('Build Database Image') {
            steps {
                echo 'Building database image...'
                dir('phase-1') {
                    script {
                        bat """
                            docker build -t ${DATABASE_IMAGE}:${IMAGE_TAG} -f db.dockerfile .
                            docker tag ${DATABASE_IMAGE}:${IMAGE_TAG} ${DATABASE_IMAGE}:latest
                        """
                    }
                }
            }
        }
        
        stage('Build Backend Image') {
            steps {
                echo 'Building backend image...'
                dir('phase-1') {
                    script {
                        bat """
                            docker build -t ${BACKEND_IMAGE}:${IMAGE_TAG} -f Dockerfile .
                            docker tag ${BACKEND_IMAGE}:${IMAGE_TAG} ${BACKEND_IMAGE}:latest
                        """
                    }
                }
            }
        }
        
        stage('Build Frontend Image') {
            steps {
                echo 'Building frontend image...'
                dir('frontend') {
                    script {
                        bat """
                            docker build --build-arg API_URL=${API_URL} -t ${FRONTEND_IMAGE}:${IMAGE_TAG} -f Dockerfile .
                            docker tag ${FRONTEND_IMAGE}:${IMAGE_TAG} ${FRONTEND_IMAGE}:latest
                        """
                    }
                }
            }
        }
        
        stage('Login to DockerHub') {
            steps {
                echo 'Logging in to DockerHub...'
                script {
                    bat 'echo %DOCKERHUB_CREDENTIALS_PSW% | docker login -u %DOCKERHUB_CREDENTIALS_USR% --password-stdin'
                }
            }
        }
        
        stage('Push Images to DockerHub') {
            steps {
                echo 'Pushing images to DockerHub...'
                script {
                    bat """
                        docker push ${DATABASE_IMAGE}:${IMAGE_TAG}
                        docker push ${DATABASE_IMAGE}:latest
                        
                        docker push ${BACKEND_IMAGE}:${IMAGE_TAG}
                        docker push ${BACKEND_IMAGE}:latest
                        
                        docker push ${FRONTEND_IMAGE}:${IMAGE_TAG}
                        docker push ${FRONTEND_IMAGE}:latest
                    """
                }
            }
        }
        
        stage('Cleanup') {
            steps {
                echo 'Cleaning up local images...'
                script {
                    bat """
                        docker rmi ${DATABASE_IMAGE}:${IMAGE_TAG} ${DATABASE_IMAGE}:latest 2>nul || echo Done
                        docker rmi ${BACKEND_IMAGE}:${IMAGE_TAG} ${BACKEND_IMAGE}:latest 2>nul || echo Done
                        docker rmi ${FRONTEND_IMAGE}:${IMAGE_TAG} ${FRONTEND_IMAGE}:latest 2>nul || echo Done
                    """
                }
            }
        }
    }
    
    post {
        always {
            echo 'Logging out from DockerHub...'
            bat 'docker logout'
        }
        success {
            echo 'Pipeline completed successfully!'
            echo "Images pushed:"
            echo "  - ${DATABASE_IMAGE}:${IMAGE_TAG} and ${DATABASE_IMAGE}:latest"
            echo "  - ${BACKEND_IMAGE}:${IMAGE_TAG} and ${BACKEND_IMAGE}:latest"
            echo "  - ${FRONTEND_IMAGE}:${IMAGE_TAG} and ${FRONTEND_IMAGE}:latest"
        }
        failure {
            echo 'Pipeline failed! Check the logs for details.'
        }
    }
}

