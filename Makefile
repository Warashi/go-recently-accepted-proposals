.DELETE_ON_ERROR:
.ONESHELL:

# --- リポジトリ設定 ---
OWNER ?= golang
REPO ?= go
LIMIT ?= 1

.PHONY: all timelines timelines-real

# デフォルトターゲット
all: timelines

# 既存のIssue番号一覧取得ルール
build/issues.txt: query.graphql
	@mkdir -p build
	gh api graphql -f query="$$(cat $<)" -F limit=$(LIMIT) -F owner=$(OWNER) -F name=$(REPO) --jq '.data.repository.issues.edges[].node.number' > $@

timelines: build/issues.txt
	@ISSUES_LIST="$$(cat build/issues.txt 2>/dev/null | tr '\r\n' ' ' | xargs)"; \
	if [ -n "$$ISSUES_LIST" ]; then \
		$(MAKE) timelines-real ISSUES="$$ISSUES_LIST"; \
	else \
		echo "No issues found in build/issues.txt"; \
	fi

# パターンルールにマッピングするための動的依存ターゲット
timelines-real: $(patsubst %, build/timeline/%.json, $(ISSUES))

# 各Issueごとの個別タイムライン取得タスク。
# ここは並列実行時 (make -j) に、それぞれ別プロセスで安全に実行されます。
build/timeline/%.json: timeline.graphql
	@mkdir -p build/timeline
	@echo "Fetching timeline for issue #$*..."
	gh api graphql --paginate -f query="$$(cat $<)" -F owner=$(OWNER) -F name=$(REPO) -F number=$* > $@
