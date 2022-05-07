FROM ubuntu:20.04

RUN ls -l

ENV TEST_ENV_1=1
ENV TEST_ENV_2=2

RUN echo $TEST_ENV_1

CMD ["sh", "-c", "echo $TEST_ENV_1 $TEST_ENV_2"]