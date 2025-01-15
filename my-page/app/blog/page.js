
        import ReactMarkdown from "react-markdown"
        import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"

    
    export default function Page(){
        return(
			<main className="flex-grow flex justify-center px-4">
				<div className="max-w-3xl w-full py-8">
				    <Card>
				    	<CardHeader>
				    	    <CardTitle>Blog Post</CardTitle>
				    	</CardHeader>
				    	<CardContent>
				    	    <article className="prose lg:prose-xl dark:prose-invert">
				    		    <ReactMarkdown>
                                   # Welcome to My Blog
                                </ReactMarkdown>
				    	    </article>
				    	</CardContent>
				    </Card>
				</div>
			</main>
        )
    }
    