#!/bin/sh -eux

cd src && npm run build:dev
cd .. && goapp serve app/web
