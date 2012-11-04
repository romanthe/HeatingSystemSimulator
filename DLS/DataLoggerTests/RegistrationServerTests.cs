using System;
using NUnit.Framework;
using DataLogger;
using System.Net.Sockets;
using NSubstitute;
using System.Text;
using System.Net;


namespace DataLoggerTests
{
	[TestFixture]
	public class RegistrationServerTests
	{
		DataCollector dataCollectorMock;
		RegistrationServer registrationServer;
		ICommunication commMock;

		[SetUp]
		public void SetUp(){
			dataCollectorMock = Substitute.For<DataCollector>();
			registrationServer = new RegistrationServer(dataCollectorMock);
			commMock = Substitute.For<ICommunication>();
		}

		[TestCase("action=register id=name","status=registration_failure")]
		[TestCase("action=register id=name port=1234","status=registered")]
		[TestCase("action=register id=name port=999999", "status=registration_failure")]
		[TestCase("action=register port=1234", "status=registration_failure")]
		public void Process_NoPortRightAnswer (string receiveMessage, string answerMessage){
			commMock.Read().Returns(receiveMessage);
			commMock.RemoteEndPoint.Returns(IPAddress.Parse("200.0.0.1"));

			registrationServer.Process(commMock);

			commMock.Received().Write(answerMessage);
		}
	}
}

