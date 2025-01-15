
		import Markdown from 'react-markdown'

	const md = `# h1 Heading

## h2 Heading

### h3 Heading

#### h4 Heading

##### h5 Heading

###### h6 Heading

blog3
`
	export default function MyApp() {
		return (
			<article class="prose lg:prose-xl">			
				<Markdown>
					{md}
				</Markdown>
			</article>
		);
	}
	