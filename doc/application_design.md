# アプリ設計

Webページに表示されるコンテンツを更新するアプリの設計方針を記載。

## 配信コンテンツ内容

配信コンテンツは以下の4つ。

- http://{ip}/index.html 
- http://{ip}/hand/rock.html
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

リリース時はストップファイルを配置してアプリ起動を抑止し
リリースディレクトリ配下にブランチのタグ情報を利用してアプリを配置する。
アプリ配置後にシンボリックリンクを切り替えることでアプリを切り替える。
アプリディレクトリは3世代まで保管する。

| パス                       | 説明               |
|----------------------------|--------------------|
| /opt/janken                | シンボリックリンク |
| /opt/release/janken/v1.0.0 | 実体               |

## アプリディレクトリ構成

| パス              | 説明               |
|-------------------|--------------------|
| /opt/janken/bin   | バイナリ           |
| /opt/janken/shell | アプリ起動用シェル |
| /opt/janken/conf  | アプリ設定ファイル |

## ログファイル

ログはrsyslogにて管理。
logrotateにより、七日以降のログは破棄。

| パス                                    | 説明   |
|-----------------------------------------|--------|
| /var/log/janken/rock/rock.log           | グー   |
| /var/log/janken/paper/paper.log         | チョキ |
| /var/log/janken/scissors/scissors.log   | パー   |

## ロックファイル

/var/run配下はOS再起動時に削除されるため
/etc/tmpfiles.d配下にディレクトリ生成の設定を追加する。

| パス                          | 説明         |
|-------------------------------|--------------|
| /var/run/janken/rock.lock     | グー         |
| /var/run/janken/paper.lock    | チョキ       |
| /var/run/janken/scissors.lock | パー         |
| /var/run/janken/stop.lock     | アプリ停止用 |
