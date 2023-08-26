import Hls from 'hls.js'
import Plyr from 'plyr'
import 'plyr/dist/plyr.css'
import { onMount } from 'solid-js'
export default function ({ source }: { source: string }) {
    let video: any
    onMount(() => {
        video.controls = true
        const defaultOptions = {}
        if (video.canPlayType('application/vnd.apple.mpegurl')) {
            video.src = source
        } else if (Hls.isSupported()) {
            const hls = new Hls()
            hls.loadSource(source)
            const player = new Plyr(video, defaultOptions)
            hls.attachMedia(video)
        } else {
            console.error(
                'This is an old browser that does not support MSE https://developer.mozilla.org/en-US/docs/Web/API/Media_Source_Extensions_API',
            )
        }
    })
    return (
        <>
            <video
                style="--plyr-color-main: #CF0031;"
                playsinline
                data-displaymaxtap
                ref={video}
            />
        </>
    )
}
