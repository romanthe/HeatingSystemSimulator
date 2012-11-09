using System;
using System.Net;
using System.Net.Sockets;
using System.Text;

namespace DataLogger
{
	public class RegistrationServer
	{
		IDataCollector m_DataCollector{get; set;}
		IPEndPoint ipep = new IPEndPoint(IPAddress.Parse("127.0.0.1"), 1236);

		public RegistrationServer (IDataCollector dataCollector)
		{
			m_DataCollector = dataCollector;
		}

		public void Process (IConnection conn)
		{
			string message = conn.Read();

			try{
				MessageParser messageParser = new MessageParser(message); 
				int port = messageParser.GetInt("port");
				string id = messageParser.GetString("id");
				m_DataCollector.AddArtifact (id, conn.RemoteIp, port);

				conn.Write("status=registered");
			}catch(Exception){
				conn.Write ("status=registration_failure");
			}
		}

		public void Start ()
		{
			TcpListener myList = new TcpListener (ipep);
			myList.Start ();
			while (true) {
				Process (new SocketConn(myList.AcceptSocket()));
			}
		}
	}

	public class SocketConn: IConnection{
		public Socket m_Socket;
		public SocketConn(Socket socket){
			m_Socket = socket;
		}

		public string Read(){
			byte[] bytes = new byte[512];
			m_Socket.Receive(bytes);
			return Encoding.ASCII.GetString(bytes);
		}

		public void Write(string message){
			m_Socket.Send(Encoding.ASCII.GetBytes(message));
		}

		public IPAddress RemoteIp{
			get{
				return (m_Socket.RemoteEndPoint as IPEndPoint).Address;
			}
		}

	}
}

