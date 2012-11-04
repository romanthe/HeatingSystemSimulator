using System;
using System.Net;
using System.Net.Sockets;
using System.Text;

namespace DataLogger
{
	public class RegistrationServer
	{
		DataCollector m_DataCollector{get; set;}
		IPEndPoint ipep = new IPEndPoint(IPAddress.Parse("127.0.0.1"), 1236);

		public RegistrationServer (DataCollector dataCollector)
		{
			m_DataCollector = dataCollector;
		}

		public void Process (ICommunication comm)
		{
			string message = comm.Read();

			try{
				MessageParser messageParser = new MessageParser(message); 
				int port = messageParser.GetInt("port");
				string id = messageParser.GetString("id");
				m_DataCollector.AddArtifact ("33", comm.RemoteEndPoint, port);

				comm.Write("status=registered");
			}catch(Exception e){
				comm.Write ("status=registration_failure");
			}
		}

		public void Start ()
		{
			TcpListener myList = new TcpListener (ipep);
			myList.Start ();
			while (true) {
				Process (new SocketComm(myList.AcceptSocket()));
			}
		}
	}

	public interface ICommunication
	{
		void Write(string message);
		string Read();
		IPAddress RemoteEndPoint{get;}
	}

	public class SocketComm: ICommunication{
		public Socket m_Socket;
		public SocketComm(Socket socket){
			m_Socket = socket;
		}

		public string Read(){
			byte[] bytes = new byte[512];
			m_Socket.Receive(bytes);
			return Encoding.ASCII.GetString(bytes);
		}

		public void Write(string message){
		//	Console.WriteLine(message);
			m_Socket.Send(Encoding.ASCII.GetBytes(message));
		}

		public IPAddress RemoteEndPoint{
			get{
				return (m_Socket.RemoteEndPoint as IPEndPoint).Address;
			}
		}

	}
}

