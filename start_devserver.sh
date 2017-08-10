#!/bin/sh -eux

cd front && npm run build:dev
cd .. && goapp serve server/src/web
