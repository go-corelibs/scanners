#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

VERSION_TAGS += SCANNERS
SCANNERS_MK_SUMMARY := go-corelibs/scanners
SCANNERS_MK_VERSION := v1.0.0

include CoreLibs.mk
