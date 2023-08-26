// @refresh reload
import { Suspense } from 'solid-js'

import {
    Body,
    ErrorBoundary,
    FileRoutes,
    Head,
    Html,
    Meta,
    Routes,
    Scripts,
    Title,
} from 'solid-start'
import 'uno.css'
import NavBar from './components/navbar'

export default function Root() {
    return (
        <Html lang="en">
            <Head>
                <Title>Torrex</Title>
                <Meta charset="utf-8" />
                <Meta
                    name="viewport"
                    content="width=device-width, initial-scale=1"
                />
            </Head>
            <Body class="flex flex-col gap-5 bg-torrex-background font-sans text-torrex-text">
                <Suspense>
                    <ErrorBoundary>
                        <NavBar />
                        <Routes>
                            <FileRoutes />
                        </Routes>
                    </ErrorBoundary>
                </Suspense>
                <Scripts />
            </Body>
        </Html>
    )
}
