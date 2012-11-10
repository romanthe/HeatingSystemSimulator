/*
 * ntp_connector.cpp
 *
 *  Created on: 10-11-2012
 *      Author: roman
 */
#ifdef linux
	#include<arpa/inet.h>
	#include<sys/socket.h>
	#include<sys/types.h>
#elif _WIN32
	#pragma comment(lib, "ws2_32.lib")
	#include<winsock2.h>
#endif

#include<string>
using namespace std;

unsigned int get_current_timestamp (string server_addr)
{
	if (server_addr == "127.127.127.127")
		return 3561569317;
	else
	{
		struct sockaddr_in server, client;

		#ifdef linux
			int sock;
		#elif _WIN32
			SOCKET sock;
			WSADATA wsaData;
			WSAStartup(MAKEWORD(2,2), &wsaData);
		#endif

		sock = socket(AF_INET,SOCK_DGRAM,IPPROTO_UDP);
		unsigned int l = sizeof(sockaddr);
		unsigned int int_sec_timestamp;
		unsigned char buffer[48] = {0xe3, 0x00, 0x06, 0xec, 0x00, 0x00, 0x00, 0x00,
									 0x00, 0x00, 0x00, 0x00, 0x49, 0x4e, 0x49, 0x54,
									 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
									 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
									 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
									 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00};
		//d4492f089ea1092d
		server.sin_family = AF_INET;
		server.sin_port = htons(123);
		server.sin_addr.s_addr = inet_addr(server_addr.c_str());

		client.sin_family = AF_INET;
		client.sin_port = htons(123);
		client.sin_addr.s_addr = INADDR_ANY;

		bind(sock, (struct sockaddr *)&client, l);

		sendto(sock, (const char *)buffer, 48, 0, (struct sockaddr *)&server, l);
		#ifdef linux
			recvfrom(sock, (char *)buffer, 48, 0, (struct sockaddr *)&client, &l);
		#elif _WIN32
			recvfrom(sock, (char *)buffer, 48, 0, (struct sockaddr *)&client, (int*)&l);
			WSACleanup();
		#endif

		int_sec_timestamp = htonl(*((unsigned int *)&buffer[40]));
		return int_sec_timestamp;
	}
}


