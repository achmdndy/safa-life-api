pipeline {
    agent any

    environment {
        // Inherits DEPLOY_HOST, DEPLOY_ROOT_DIR from Jenkins Global Config
        SERVICE_NAME = "safalife-core"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        // 'Build' is handled during Deploy phase on the VPS to simplify setup 
        // and avoid Docker-in-Docker context issues.


        stage('Deploy') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: 'vps-login', usernameVariable: 'VPS_USER', passwordVariable: 'VPS_PASS')]) {
                         sh """
                            sshpass -p '\$VPS_PASS' ssh -o StrictHostKeyChecking=no \$VPS_USER@\${DEPLOY_HOST##*@} << EOF
                                cd ${DEPLOY_ROOT_DIR}/core
                                git pull origin main
                                # Rebuild/Restart
                                docker compose up -d --build --no-deps ${SERVICE_NAME}-api
                            EOF
                        """
                    }
                }
            }
        }
    }
}
