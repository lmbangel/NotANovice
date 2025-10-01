import grpc
from concurrent import futures
import quiz_pb2
import quiz_pb2_grpc


class QuizAgentService(quiz_pb2_grpc.QuizAgentServiceServicer):
    def GetQuizOfTheDay(self, request, context):
        return quiz_pb2.Quiz(
            quiz_id=1,
            id=2,
            question="Question",
            correct_answer="A",
            timestamp=1000,
            is_active=1,
            a_answer="A",
            b_answer="B",
            c_answer="C",
            d_answer="D"
        )


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    quiz_pb2_grpc.add_QuizAgentServiceServicer_to_server(QuizAgentService(), server)
    server.add_insecure_port("[::]:50052")
    server.start()
    print("Python gRPC server running on :50052")

    server.wait_for_termination()


if __name__ == "__main__":
    serve()
