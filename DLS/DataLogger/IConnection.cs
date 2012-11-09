using System;
using System.Net;
using System.Net.Sockets;
using System.Text;

namespace DataLogger
{
	public interface IConnection
	{
		void Write(string message);
		string Read();
		IPAddress RemoteIp{get;}
	}
}

