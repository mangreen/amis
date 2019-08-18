root:              .
tmp_path:          ./tmp
build_name:        runner-build
build_log:         runner-build-errors.log
valid_ext:         .go, .tpl, .tmpl, .html
no_rebuild_ext:    .tpl, .tmpl, .html
ignored:           assets, tmp, configs, docker, k8s, node_modules, migrations, public, views, web
build_delay:       600
colors:            1
log_color_main:    cyan
log_color_build:   yellow
log_color_runner:  green
log_color_watcher: magenta
log_color_app:
