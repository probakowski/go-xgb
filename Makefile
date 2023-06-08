# This Makefile is used by the developer. It is not needed in any way to build
# a checkout of the XGB repository.
# It will be useful, however, if you are hacking at the code generator.
# i.e., after making a change to the code generator, run 'make' in the
# xgb directory. This will build xgbgen and regenerate each sub-package.
# 'make test' will then run any appropriate tests (just tests xproto right now).
# 'make bench' will test a couple of benchmarks.
# 'make build-all' will then try to build each extension. This isn't strictly
# necessary, but it's a good idea to make sure each sub-package is a valid
# Go package.

# My path to the X protocol XML descriptions.
ifndef XPROTO
XPROTO=/usr/share/xcb
endif

# All of the XML files in my /usr/share/xcb directory
# This is intended to build xgbgen and generate Go code for each supported
# extension. IGNORE: xkb.xml, glx.xml
all: \
		 bigreq.xml composite.xml damage.xml dpms.xml dri2.xml \
		 ge.xml randr.xml record.xml res.xml \
		 render.xml screensaver.xml shape.xml shm.xml xc_misc.xml \
		 xevie.xml xf86dri.xml xf86vidmode.xml xfixes.xml xinerama.xml \
		 xprint.xml xproto.xml xselinux.xml xtest.xml \
		 xvmc.xml xv.xml

# Builds each individual sub-package to make sure its valid Go code.
build-all: bigreq.b composite.b damage.b dpms.b dri2.b ge.b glx.b randr.b \
					 record.b render.b res.b screensaver.b shape.b shm.b xcmisc.b \
					 xevie.b xf86dri.b xf86vidmode.b xfixes.b xinerama.b \
					 xprint.b xproto.b xselinux.b xtest.b xv.b xvmc.b

%.b:
	(cd $* ; go build)

# xc_misc is special because it has an underscore.
# There's probably a way to do this better, but Makefiles aren't my strong suit.
xc_misc.xml:
	mkdir -p xcmisc
	go run ./cmd/xgbgen --gofmt=true --proto-path $(XPROTO) $(XPROTO)/xc_misc.xml > xcmisc/xcmisc.go

%.xml:
	mkdir -p $*
	go run ./cmd/xgbgen --gofmt=true --proto-path $(XPROTO) $(XPROTO)/$*.xml > $*/$*.go

# Just test the xproto core protocol for now.
test:
	(cd xproto ; go test)

# Force all xproto benchmarks to run and no tests.
bench:
	(cd xproto ; go test -run 'nomatch' -bench '.*' -cpu 1,2,3,6)

