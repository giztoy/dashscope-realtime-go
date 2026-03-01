# text-chat 示例说明（DashScope Realtime 文本能力）

这个示例用于验证 **DashScope Realtime 的文本输出能力**，并明确记录当前已知限制。

## 关键结论

1. **可以输出文本**
   - 通过 `response.create` 触发，服务端会返回 `response.text.delta` / `choices` 等文本事件。
2. **纯文本输入通道不稳定（重点）**
   - `input_text_buffer.append` 在真实环境下可能返回参数错误（例如 `invalid_value`）。
   - 因此本示例默认不依赖该路径，而是采用 `response.create` + `instructions` 的稳定方案。

> 结论：**可以做“文本回复型”会话；不建议把它当作“稳定纯文本输入聊天通道”。**

## 运行方式

```bash
go run ./examples/text-chat -rounds 3 -prompt "你好，请用一句话自我介绍"
```

必需环境变量：

- `DASHSCOPE_API_KEY`

可选环境变量：

- `DASHSCOPE_MODEL`
- `DASHSCOPE_BASE_URL`
- `DASHSCOPE_PRO_MODEL`

## 重要参数

- `-rounds`: 多轮对话轮数（默认 3）
- `-test-pro-settings`: 探测运行时 `session.update` 设置切换（默认 true）
- `-test-history-edit`: 探测历史改写能力（默认 true）
- `-probe-append-text`: 额外探测 `AppendText` 路径（可能失败，默认 false）
- `-probe-cancel`: 额外探测 `CancelResponse` 路径（默认 false）
- `-demo-error`: 演示可复现错误路径（本地参数错误，默认 false）

## 为什么有 `history_rewrite=false`

示例里会用 `SendRaw + response.create(messages)` 做“历史改写能力探测”。

- 若返回能明确体现改写结果，则标记为 `true`
- 若未体现或接口不支持该语义，则标记为 `false`

当前实测里经常是 `false`，这说明在该接口/模型上，这条能力并不稳定或不保证支持。

## 错误演示（可复现）

```bash
go run ./examples/text-chat -demo-error
```

会演示并打印：

- `AppendText("")` 的本地参数错误（`InvalidParameter`）
- `SendRaw(empty)` 的本地参数错误（`InvalidParameter`）
