pipeline {
    agent any
    triggers { pollSCM('H/2 * * * *') }
    options {
        timeout(time: 20, unit: 'MINUTES')
    }
    stages {
        stage('Build') {
            steps {
                script {
                    openshift.withCluster() {
                        openshift.withProject() {
                            echo "Using project: ${openshift.project()}"
                            def bc = openshift.selector( 'buildconfig/next-build' )
                            def buildSelector = bc.startBuild()
                            buildSelector.logs('-f')
                        }
                    }
                }
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}