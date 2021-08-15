type CheckMessage = {
    name: string
    duration: number
    is_success: boolean
    checked_at: Date
    result: any
}
type CheckMeta = {
    name: string
    description: string
    interval: string
}

export type { CheckMessage, CheckMeta }
