
		import Link from 'next/link'
		import { Button } from "@/components/ui/button"

	
		export default function Header() {
			return (
				<header className="border-b">
					<div className="container mx-auto max-w-3xl py-4">
						<nav className="flex justify-between items-center">
						<Link href="/" className="text-2xl font-bold">
							My Blog
						</Link>
						<div className="flex space-x-2">
							<Button variant="ghost" asChild>
							<Link href="/">Home</Link>
							</Button>
							<Button variant="ghost" asChild>
							<Link href="/about">About</Link>
							</Button>
							<Button variant="ghost" asChild>
							<Link href="/contact">Contact</Link>
							</Button>
						</div>
						</nav>
					</div>
				</header>
			)
		}
	