#!/bin/sh
if [ -z "${VERBOSE}" ]; then
  /app/xdev server --port "${PORT}" 
else
  /app/xdev server --port "${PORT}" --verbose
fi