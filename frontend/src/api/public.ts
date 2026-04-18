/**
 * Public (unauthenticated) API endpoints used by the marketing homepage.
 *
 * These endpoints never require a logged-in user and are safe to call before
 * the auth state is hydrated.
 */

import { apiClient } from './client'

export interface PublicPlan {
  id: number
  name: string
  description: string
  price: number
  original_price?: number | null
  validity_days: number
  validity_unit: string
  features: string[]
  product_name: string
  sort_order: number
  group_id: number
  group_name: string
  group_platform: string
  rate_multiplier: number
  daily_limit_usd?: number | null
  weekly_limit_usd?: number | null
  monthly_limit_usd?: number | null
  supported_model_scopes: string[]
}

export type PlatformHealthStatus = 'healthy' | 'degraded' | 'down' | 'unknown'

export interface PublicHealthItem {
  platform: string
  display_name: string
  total_accounts: number
  available_count: number
  rate_limit_count: number
  error_count: number
  status: PlatformHealthStatus
}

export interface PublicHealthResponse {
  enabled: boolean
  collected_at?: string
  platforms: PublicHealthItem[]
}

export type ServerLineStatus = 'healthy' | 'degraded' | 'down' | 'unknown'

export interface PublicServerLine {
  id: string
  name: string
  region: string
  url: string
  probe_path: string
  description: string
  sort_order: number
  enabled: boolean
  status: ServerLineStatus
  latency_ms: number
  checked_at: string
  error?: string
}

export const publicAPI = {
  /** List subscription plans for the marketing homepage (no auth required). */
  getPlans() {
    return apiClient.get<PublicPlan[]>('/public/plans')
  },

  /** Aggregated platform-level health snapshot. */
  getModelHealth() {
    return apiClient.get<PublicHealthResponse>('/public/model-health')
  },

  /** Enabled server lines enriched with probe status. */
  getServerLines() {
    return apiClient.get<PublicServerLine[]>('/public/server-lines')
  }
}
