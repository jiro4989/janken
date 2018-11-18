# アプリ設計

Webページに表示されるコンテンツを更新するアプリの設計方針を記載。

## 配信コンテンツ内容

配信コンテンツは以下の4つ。

- http://{ip}/index.html http://{ip}/hand/rock.html
- http://{ip}/hand/paper.html
- http://{ip}/hand/scissors.html

コンテンツの内容は下記の通り。

### index.html

以下のリンクを記載。

- rock.htmlのリンク
- paper.htmlのリンク
- scissors.htmlのリンク

### rock.html

グーを出したときの相手(Webシステム)側の出した手と、その勝敗。

### paper.html

チョキを出したときの相手(Webシステム)側の出した手と、その勝敗。

### scissors.html

パーを出したときの相手(Webシステム)側の出した手と、その勝敗。

## アプリ処理概要

前述のコンテンツを生成するためにGoアプリを作成し、Webサーバ上に配置する。
後述のトリガーによりGoアプリを起動し、コンテンツを更新する。

### アプリ実行トリガー

毎分cronを起動し、1分の起動の間に1秒おきにGoプロセスを起動し、htmlコンテンツを更
新する。

### アプリ処理ロジック

![アクティビティ図](./img/application.png)

## アプリ配置先

| パス                       | 説明               |
|----------------------------|--------------------|
| /opt/janken                | シンボリックリンク |
| /opt/release/janken/v1.0.0 | 実体               |

## ログファイル

ログはrsyslogにて管理。
logrotateにより、七日以降のログは破棄。

- /var/log/janken/rock/rock.log
- /var/log/janken/paper/paper.log
- /var/log/janken/scissors/scissors.log

## ロックファイル

以下の通りロックファイルを生成する。

- /var/run/janken/rock.lock
- /var/run/janken/paper.lock
- /var/run/janken/scissors.lock

/var/run配下はOS再起動時に削除されるため
/etc/tmpfiles.d配下にディレクトリ生成の設定を追加する。
