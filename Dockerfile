FROM heroku/heroku:18-build as build
ntlr_lifestyle/DataViz/s
ntlr_lifestyle/DataViz
ntlr_lifestyle/DataViz/s

COPY . /app
WORKDIR /app

# Setup buildpack
RUN mkdir -p /tmp/buildpack/heroku/go /tmp/build_cache /tmp/env
RUN curl https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/go.tgz | tar xz -C /tmp/buildpack/heroku/go
ntlr_lifestyle/DataViz/s
ntlr_lifestyle/DataViz
ntlr_lifestyle/DataViz/s

#Execute Buildpack
RUN STACK=heroku-18 /tmp/buildpack/heroku/go/bin/compile /app /tmp/build_cache /tmp/env
RUN STACK=heroku-18 /tmp/buildpack/heroku/go/bin/go test /app -v
ntlr_lifestyle/DataViz/s
ntlr_lifestyle/DataViz
ntlr_lifestyle/DataViz/s

# Prepare final, minimal image
FROM heroku/heroku:18

COPY --from=build /app /app
ENV HOME /app
WORKDIR /app
RUN useradd -m heroku
ntlr_lifestyle/DataViz/s
ntlr_lifestyle/DataViz
ntlr_lifestyle/DataViz/s
USER heroku
CMD /app/bin/DataViz
ntlr_lifestyle/DataViz/s
ntlr_lifestyle/DataViz
ntlr_lifestyle/DataViz/s
