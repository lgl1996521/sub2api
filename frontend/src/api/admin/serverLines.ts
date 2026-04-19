/**
 * Admin Server Lines API endpoints
 * Manages the user-facing server line list shown on the marketing homepage.
 */

import { apiClient } from '../client'

export interface ServerLine {
  id: string
  name: string
  region: string
  url: string
  probe_path: string
  description: string
  sort_order: number
  enabled: boolean
}

export interface ServerLineWithStatus extends ServerLine {
  status: 'healthy' | 'degraded' | 'down' | 'unknown'
  latency_ms: number
  checked_at: string
  error?: string
}

const adminServerLinesAPI = {
  list() {
    return apiClient.get<ServerLine[]>('/admin/server-lines')
  },
  listWithStatus() {
    return apiClient.get<ServerLineWithStatus[]>('/admin/server-lines/status')
  },
  save(lines: ServerLine[]) {
    return apiClient.put<ServerLine[]>('/admin/server-lines', { lines })
  }
}

export default adminServerLinesAPI
