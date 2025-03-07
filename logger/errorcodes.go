package logger

type errorcode int

// Validation Error Codes
const (
	// 1001: Required field missing in the API request.
	ErrReqFieldMissing errorcode = 1001
	// 1002: Invalid data format or type provided.
	ErrInvalidData errorcode = 1002
	// 1003: Value exceeds allowed range.
	ErrValueExceedsRange errorcode = 1003
	// 1004: Unsupported file type uploaded.
	ErrUnsupportedFile errorcode = 1004
	// 1005: Duplicate data found in the request.
	ErrDuplicateData errorcode = 1005
	// 1006: Invalid query parameter or filter provided.
	ErrInvalidQuery errorcode = 1006
	// 1007: CSRF token validation failed.
	ErrCSRFTokenInvalid errorcode = 1007
	// 1008: File size exceeds allowed limit.
	ErrFileSizeExceeded errorcode = 1008
)

// Authentication Error Codes
const (
	// 2001: User is not authenticated.
	ErrNotAuthenticated errorcode = 2001
	// 2002: User does not have permission to access this resource.
	ErrPermissionDenied errorcode = 2002
	// 2003: Invalid or expired authentication token.
	ErrInvalidToken errorcode = 2003
	// 2004: Account temporarily locked due to multiple failed attempts.
	ErrAccountLocked errorcode = 2004
	// 2005: Session expired; re-authentication required.
	ErrSessionExpired errorcode = 2005
	// 2006: Multi-factor authentication required.
	ErrMFARequired errorcode = 2006
	// 2007: Invalid OAuth token.
	ErrInvalidOAuthToken errorcode = 2007
)

// Resource Error Codes
const (
	// 3001: Resource not found (e.g., product, order).
	ErrResourceNotFound errorcode = 3001
	// 3002: Resource is currently unavailable or locked.
	ErrResourceLocked errorcode = 3002
	// 3003: Insufficient inventory for requested product.
	ErrInsufficientInventory errorcode = 3003
	// 3004: Resource has been archived or deleted.
	ErrResourceArchived errorcode = 3004
	// 3005: Dependency not found (e.g., related resource missing).
	ErrDependencyNotFound errorcode = 3005
	// 3006: Conflict detected in resource update.
	ErrResourceConflict errorcode = 3006
	// 3007: Read-only resource modification attempted.
	ErrReadOnlyResource errorcode = 3007
)

// System Error Codes
const (
	// 4001: Internal server error.
	ErrInternalServer errorcode = 4001
	// 4002: Service is temporarily unavailable.
	ErrServiceUnavailable errorcode = 4002
	// 4003: Database connection error.
	ErrDatabaseError errorcode = 4003
	// 4004: Cache synchronization failed.
	ErrCacheSyncFailed errorcode = 4004
	// 4005: Unexpected behavior in background job processing.
	ErrJobProcessingError errorcode = 4005
	// 4006: Memory usage exceeded safe threshold.
	ErrHighMemoryUsage errorcode = 4006
	// 4007: Disk space running low.
	ErrLowDiskSpace errorcode = 4007
)

// Integration Error Codes
const (
	// 5001: Third-party API returned an error.
	ErrAPIError errorcode = 5001
	// 5002: Failed to connect to an external service.
	ErrConnectionFailed errorcode = 5002
	// 5003: Timeout while waiting for a third-party API response.
	ErrAPITimeout errorcode = 5003
	// 5004: Invalid response received from third-party service.
	ErrInvalidAPIResponse errorcode = 5004
	// 5005: API quota limit reached for external service.
	ErrAPILimitReached errorcode = 5005
	// 5006: Webhook delivery failed.
	ErrWebhookFailed errorcode = 5006
	// 5007: External service returned an authentication error.
	ErrExternalAuthError errorcode = 5007
)

// Business Logic Error Codes
const (
	// 6001: Order cannot be processed due to invalid status.
	ErrInvalidOrderStatus errorcode = 6001
	// 6002: Merchant quota exceeded for daily requests.
	ErrMerchantQuotaExceeded errorcode = 6002
	// 6003: Payment gateway rejected the transaction.
	ErrPaymentRejected errorcode = 6003
	// 6004: Refund cannot be processed due to insufficient balance.
	ErrRefundFailed errorcode = 6004
	// 6005: Promotion code is invalid or expired.
	ErrInvalidPromoCode errorcode = 6005
	// 6006: Order cancellation window has passed.
	ErrCancellationWindowClosed errorcode = 6006
	// 6007: Subscription plan limit reached.
	ErrSubscriptionLimitReached errorcode = 6007
	// 6008: Cannot modify order after fulfillment.
	ErrOrderModificationNotAllowed errorcode = 6008
)

// Info Logs (7000 - 7499)
const (
	// 7001: User successfully authenticated.
	InfoUserAuthenticated errorcode = 7001
	// 7002: Cache hit for requested resource.
	InfoCacheHit errorcode = 7002
	// 7003: Request processed successfully.
	InfoRequestProcessed errorcode = 7003
	// 7004: Background job completed successfully.
	InfoJobCompleted errorcode = 7004
	// 7005: External API request completed successfully.
	InfoExternalAPIRequestSuccess errorcode = 7005
)

// Warning Logs (7500 - 7999)
const (
	// 7501: High response time detected.
	WarnHighResponseTime errorcode = 7501
	// 7502: Deprecated API version used in request.
	WarnDeprecatedAPIVersion errorcode = 7502
	// 7503: Soft limit exceeded for resource usage.
	WarnSoftLimitExceeded errorcode = 7503
	// 7504: Retryable error occurred in background job.
	WarnJobRetryableError errorcode = 7504
	// 7505: External API returned a warning.
	WarnExternalAPIWarning errorcode = 7505
)
