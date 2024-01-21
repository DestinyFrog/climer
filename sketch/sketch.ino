#include <DHT_U.h>
#include <DHT.h>
#include <SPI.h>
#include <UIPEthernet.h>

byte mac[] = { 0xDE, 0xAD, 0xBE, 0xEF, 0xFE, 0xED };
byte ip[] = { 192, 168, 1, 175 };
EthernetServer server(80);

DHT dht = DHT( A0, DHT11 );

void setup(){
	Ethernet.begin(mac,ip);
	server.begin();
	dht.begin();
}

void loop(){
	float h = dht.readHumidity();
	float t = dht.readTemperature();
	EthernetClient client = server.available();

	if(client)
		do {
			if ( client.available() ) {
				// Monta Requisicao http
				client.println("HTTP/1.1 200 OK");
				client.println("Content-Type: application/json");
				client.println();
				client.print( "{\"umidade\":" );
				client.print( h );
				client.print( ",\"temperatura\":" );
				client.print( t );
				client.println( '}' );
				client.stop();
			}
		} while( client.connected() );
}
