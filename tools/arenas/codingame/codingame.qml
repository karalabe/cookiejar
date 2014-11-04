import QtQuick 2.1
import QtWebEngine 1.0
import QtQuick.Controls 1.0

ApplicationWindow {
    width: 1024
    height: 768
    
    WebEngineView {
    	objectName: "arena"
        anchors.fill: parent
        url: "http://codingame.com"
    }
}