pipeline {
    agent any
    stages {
        environment {
            def TEST_SCRIPT = "test_HW1.py"
            def PORT_EXT = "8080"
            def PORT_INT = "8080"
            def CONTAINER = "CS128-HW1"
            def TAG = "local:${CONTAINER}"
            def BUILD_FLAGS = "--force-rm --no-cache"
            def RUN_FLAGS = "-p ${PORT_EXT}:${PORT_INT} -d --name ${CONTAINER} --rm ${TAG}"
        }
        stage('Build') {
            steps {
                echo 'Building container...'
                sh "docker build ${BUILD_FLAGS} ${TAG} ."
            }
        }
        stage('Run') {
            steps {
                echo 'Running container...'
                sh "docker run ${RUN_FLAGS}"
            }
        }
        stage('Test') {
            steps {
                echo 'Testing app...'
                sh "./${TEST_SCRIPT}"
            }
        }
        stage('Cleanup') {
            steps {
                echo 'Cleaning up...'
                sh "docker stop ${CONTAINER}"
            }
        }
    }
}
