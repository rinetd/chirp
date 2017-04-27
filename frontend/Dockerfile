FROM node:7.6
# for unbuntu debian
RUN echo "Asia/Shanghai" > /etc/timezone
RUN dpkg-reconfigure -f noninteractive tzdata
RUN sed -i "s|deb.debian.org|mirrors.aliyun.com|g" /etc/apt/sources.list

# for China
ENV NODEJS_ORG_MIRROR="http://npm.taobao.org/mirrors/node"
ENV SASS_BINARY_SITE="https://npm.taobao.org/mirrors/node-sass"
ENV PHANTOMJS_CDNURL="http://npm.taobao.org/mirrors/phantomjs"
ENV ELECTRON_MIRROR="https://npm.taobao.org/mirrors/electron/"
ENV CHROMEDRIVER_CDNURL="http://npm.taobao.org/mirrors/chromedriver"

# npm set registry https://registry.npm.taobao.org # 注册模块镜像
# npm set disturl https://npm.taobao.org/dist # node-gyp 编译依赖的 node 源码镜像
#
# ## 以下选择添加
# npm set chromedriver_cdnurl http://cdn.npm.taobao.org/dist/chromedriver # chromedriver 二进制包镜像
# npm set operadriver_cdnurl http://cdn.npm.taobao.org/dist/operadriver # operadriver 二进制包镜像
# npm set phantomjs_cdnurl http://cdn.npm.taobao.org/dist/phantomjs # phantomjs 二进制包镜像
# npm set sass_binary_site http://cdn.npm.taobao.org/dist/node-sass # node-sass 二进制包镜像
# npm set electron_mirror http://cdn.npm.taobao.org/dist/electron/ # electron 二进制包镜像
#
# npm cache clean # 清空缓存


# RUN npm config set registry https://registry.npm.taobao.org \
# 淘宝 NPM 镜像
# && npm install -g cnpm --registry=https://registry.npm.taobao.org \
  # 安装 Yarn
# && cnpm install -g yarn


# Fix bug https://github.com/npm/npm/issues/9863
# RUN cd $(npm root -g)/npm \
#     && npm install fs-extra --registry=https://registry.npm.taobao.org \
#     && sed -i -e s/graceful-fs/fs-extra/ -e s/fs\.rename/fs.move/ ./lib/utils/rename.js

# RUN npm install yarn -g \
#     && yarn config set registry https://registry.npm.taobao.org
# install yarn
# RUN \
#   curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - && \
#   echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list && \
#   apt-get update && apt-get install yarn

ENV NODE_ENV=development
RUN yarn config set registry https://registry.npm.taobao.org

# install global dependencies
RUN yarn global add typescript webpack webpack-dev-server tslint typings

# install all packes
RUN mkdir /frontend
COPY package.json /frontend/
WORKDIR /frontend/

RUN yarn install

# copy other files
COPY . /frontend
