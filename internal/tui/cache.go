package tui

import (
	"sync"
	"time"

	"github.com/sebastiengiband/dotfiles/internal/scripts"
)

// CacheEntry represents a cached item
type CacheEntry struct {
	Value     interface{}
	ExpiresAt time.Time
	CreatedAt time.Time
}

// IsExpired checks if the cache entry has expired
func (ce CacheEntry) IsExpired() bool {
	return time.Now().After(ce.ExpiresAt)
}

// Cache represents a thread-safe cache with TTL
type Cache struct {
	items map[string]CacheEntry
	mutex sync.RWMutex
	ttl   time.Duration
}

// NewCache creates a new cache with the specified TTL
func NewCache(ttl time.Duration) *Cache {
	cache := &Cache{
		items: make(map[string]CacheEntry),
		ttl:   ttl,
	}

	// Start cleanup goroutine
	go cache.cleanup()

	return cache
}

// Set stores a value in the cache
func (c *Cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	c.items[key] = CacheEntry{
		Value:     value,
		ExpiresAt: now.Add(c.ttl),
		CreatedAt: now,
	}
}

// Get retrieves a value from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, exists := c.items[key]
	if !exists || entry.IsExpired() {
		return nil, false
	}

	return entry.Value, true
}

// Delete removes a value from the cache
func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.items, key)
}

// Clear removes all items from the cache
func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.items = make(map[string]CacheEntry)
}

// Size returns the number of items in the cache
func (c *Cache) Size() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return len(c.items)
}

// cleanup removes expired entries periodically
func (c *Cache) cleanup() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.items {
			if entry.IsExpired() {
				delete(c.items, key)
			}
		}
		c.mutex.Unlock()
	}
}

// CachedScriptRunner wraps ScriptRunner with caching
type CachedScriptRunner struct {
	runner *scripts.ScriptRunner
	cache  *Cache
}

// NewCachedScriptRunner creates a new cached script runner
func NewCachedScriptRunner(runner *scripts.ScriptRunner) *CachedScriptRunner {
	return &CachedScriptRunner{
		runner: runner,
		cache:  NewCache(5 * time.Minute), // 5 minute TTL
	}
}

// CheckCommand checks if a command exists (cached)
func (csr *CachedScriptRunner) CheckCommand(cmd string) bool {
	cacheKey := "cmd_" + cmd

	if cached, found := csr.cache.Get(cacheKey); found {
		return cached.(bool)
	}

	result := csr.runner.CheckCommand(cmd)
	csr.cache.Set(cacheKey, result)

	return result
}

// GetSystemInfo returns system information (cached)
func (csr *CachedScriptRunner) GetSystemInfo() map[string]string {
	cacheKey := "system_info"

	if cached, found := csr.cache.Get(cacheKey); found {
		return cached.(map[string]string)
	}

	result := csr.runner.GetSystemInfo()
	csr.cache.Set(cacheKey, result)

	return result
}

// GetInstalledTools returns installed tools (cached)
func (csr *CachedScriptRunner) GetInstalledTools() map[string]bool {
	cacheKey := "installed_tools"

	if cached, found := csr.cache.Get(cacheKey); found {
		return cached.(map[string]bool)
	}

	result := csr.runner.GetInstalledTools()
	csr.cache.Set(cacheKey, result)

	return result
}

// GetPackageManager returns the package manager (cached)
func (csr *CachedScriptRunner) GetPackageManager() string {
	cacheKey := "package_manager"

	if cached, found := csr.cache.Get(cacheKey); found {
		return cached.(string)
	}

	result := csr.runner.GetPackageManager()
	csr.cache.Set(cacheKey, result)

	return result
}

// InvalidateCache clears all cached data
func (csr *CachedScriptRunner) InvalidateCache() {
	csr.cache.Clear()
}

// FileContentCache caches file contents for the editor
type FileContentCache struct {
	cache *Cache
}

// NewFileContentCache creates a new file content cache
func NewFileContentCache() *FileContentCache {
	return &FileContentCache{
		cache: NewCache(2 * time.Minute), // 2 minute TTL for file contents
	}
}

// GetFileContent gets cached file content
func (fcc *FileContentCache) GetFileContent(filePath string) (string, bool) {
	if cached, found := fcc.cache.Get(filePath); found {
		return cached.(string), true
	}
	return "", false
}

// SetFileContent caches file content
func (fcc *FileContentCache) SetFileContent(filePath, content string) {
	fcc.cache.Set(filePath, content)
}

// InvalidateFile removes a file from cache
func (fcc *FileContentCache) InvalidateFile(filePath string) {
	fcc.cache.Delete(filePath)
}

// CacheManager manages all caches in the application
type CacheManager struct {
	scriptCache *CachedScriptRunner
	fileCache   *FileContentCache
}

// NewCacheManager creates a new cache manager
func NewCacheManager(scriptRunner *scripts.ScriptRunner) *CacheManager {
	return &CacheManager{
		scriptCache: NewCachedScriptRunner(scriptRunner),
		fileCache:   NewFileContentCache(),
	}
}

// GetScriptRunner returns the cached script runner
func (cm *CacheManager) GetScriptRunner() *CachedScriptRunner {
	return cm.scriptCache
}

// GetFileCache returns the file content cache
func (cm *CacheManager) GetFileCache() *FileContentCache {
	return cm.fileCache
}

// InvalidateAll clears all caches
func (cm *CacheManager) InvalidateAll() {
	cm.scriptCache.InvalidateCache()
	cm.fileCache.cache.Clear()
}

// GetCacheStats returns cache statistics
func (cm *CacheManager) GetCacheStats() map[string]int {
	return map[string]int{
		"script_cache_size": cm.scriptCache.cache.Size(),
		"file_cache_size":   cm.fileCache.cache.Size(),
	}
}

// Global cache manager instance
var globalCacheManager *CacheManager

// InitializeCache initializes the global cache manager
func InitializeCache(scriptRunner *scripts.ScriptRunner) {
	globalCacheManager = NewCacheManager(scriptRunner)
}

// GetCacheManager returns the global cache manager
func GetCacheManager() *CacheManager {
	return globalCacheManager
}
