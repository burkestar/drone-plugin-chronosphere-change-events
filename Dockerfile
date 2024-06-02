FROM python:3.12-alpine

ENV PYTHONUNBUFFERED=1

RUN pip install --no-cache-dir 'requests~=2.32.3'

COPY cli /usr/local/bin

ENTRYPOINT [ "cli" ]
