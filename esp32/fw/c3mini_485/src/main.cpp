#include <Arduino.h>
#include <ArduinoOTA.h>


void setup() {

  ArduinoOTA
    .onStart([]() {
      String type;
      if (ArduinoOTA.getCommand() == U_FLASH) {
        type = "sketch";
      } else {  // U_SPIFFS
        type = "filesystem";
      }

      // NOTE: if updating SPIFFS this would be the place to unmount SPIFFS using SPIFFS.end()
   })
    .onEnd([]() {
    })
    .onProgress([](unsigned int progress, unsigned int total) {
    })
    .onError([](ota_error_t error) {
      if (error == OTA_AUTH_ERROR) {
        
      } else if (error == OTA_BEGIN_ERROR) {
        
      } else if (error == OTA_CONNECT_ERROR) {
        
      } else if (error == OTA_RECEIVE_ERROR) {
        
      } else if (error == OTA_END_ERROR) {
        
      }
    });


}

void loop() {
    ArduinoOTA.handle();
}
