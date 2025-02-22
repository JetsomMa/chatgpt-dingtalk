package process

import "strings"

func GeneratePrompt(msg string) string {
	switch {
	case strings.HasPrefix(msg, "#周报"):
		return "请帮我把以下的工作内容填充为一篇完整的周报，用 markdown 格式以分点叙述的形式输出：" + strings.Replace(msg, "#周报", "", -1)
	case strings.HasPrefix(msg, "#前端"):
		return `我想让你充当前端开发专家。我将提供一些关于 Js、Node 等前端代码问题的具体信息，而你的工作就是想出为我解决问题的策略。这可能包括建议代码、代码逻辑思路策略。我的第一个请求是：` + strings.Replace(msg, "#前端", "", -1)
	case strings.HasPrefix(msg, "#架构师"):
		return `我希望你担任 IT 架构师。我将提供有关应用程序或其他数字产品功能的一些详细信息，而您的工作是想出将其集成到 IT 环境中的方法。这可能涉及分析业务需求、执行差距分析以及将新系统的功能映射到现有 IT 环境。接下来的步骤是创建解决方案设计、物理网络蓝图、系统集成接口定义和部署环境蓝图。我的第一个请求是：` + strings.Replace(msg, "#架构师", "", -1)
	case strings.HasPrefix(msg, "#产品经理"):
		return `请确认我的以下请求。请您作为产品经理回复我。我将会提供一个主题，您将帮助我编写一份包括以下章节标题的 PRD 文档：主题、简介、问题陈述、目标与目的、用户故事、技术要求、收益、KPI 指标、开发风险以及结论。在我要求具体主题、功能或开发的 PRD 之前，请不要先写任何一份 PRD 文档。` + strings.Replace(msg, "#产品经理", "", -1)
	case strings.HasPrefix(msg, "#网络安全"):
		return `我想让你充当网络安全专家。我将提供一些关于如何存储和共享数据的具体信息，而你的工作就是想出保护这些数据免受恶意行为者攻击的策略。这可能包括建议加密方法、创建防火墙或实施将某些活动标记为可疑的策略。我的第一个请求是：` + strings.Replace(msg, "#网络安全", "", -1)
	case strings.HasPrefix(msg, "#正则"):
		return `我希望你充当正则表达式生成器。您的角色是生成匹配文本中特定模式的正则表达式。您应该以一种可以轻松复制并粘贴到支持正则表达式的文本编辑器或编程语言中的格式提供正则表达式。不要写正则表达式如何工作的解释或例子；只需提供正则表达式本身。我的第一个提示是：：` + strings.Replace(msg, "#正则", "", -1)
	case strings.HasPrefix(msg, "#招聘"):
		return `我想让你担任招聘人员。我将提供一些关于职位空缺的信息，而你的工作是制定寻找合格申请人的策略。这可能包括通过社交媒体、社交活动甚至参加招聘会接触潜在候选人，以便为每个职位找到最合适的人选。我的第一个请求是：` + strings.Replace(msg, "#招聘", "", -1)
	case strings.HasPrefix(msg, "#知乎"):
		return `知乎的风格是：用"谢邀"开头，用很多学术语言，引用很多名言，做大道理的论述，提到自己很厉害的教育背景并且经验丰富，最后还要引用一些论文。请用知乎风格：` + strings.Replace(msg, "#知乎", "", -1)
	case strings.HasPrefix(msg, "#翻译"):
		return `下面我让你来充当翻译家，你的目标是把任何语言翻译成中文，请翻译时不要带翻译腔，而是要翻译得自然、流畅和地道，最重要的是要简明扼要。请翻译下面这句话：` + strings.Replace(msg, "#翻译", "", -1)
	case strings.HasPrefix(msg, "#小红书"):
		return "小红书的风格是：很吸引眼球的标题，每个段落都加 emoji, 最后加一些 tag。请用小红书风格：" + strings.Replace(msg, "#小红书", "", -1)
	case strings.HasPrefix(msg, "#解梦"):
		return "我要你充当解梦师。我会给你描述我的梦，你会根据梦中出现的符号和主题提供解释。不要提供关于梦者的个人意见或假设。仅根据所提供的信息提供事实解释。我的第一个梦是：" + strings.Replace(msg, "#解梦", "", -1)
	default:
		return msg
	}
}
