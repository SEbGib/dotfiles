# MCP Security Documentation

## Overview

This document outlines the security framework implemented for Model Context Protocol (MCP) servers in this dotfiles configuration. The security-first approach ensures safe integration of MCP servers while maintaining functionality.

## Security Architecture

### 🔒 Core Security Principles

1. **Zero Trust**: Every MCP server is assessed before installation
2. **Principle of Least Privilege**: Minimal required permissions
3. **Defense in Depth**: Multiple security layers
4. **Continuous Monitoring**: Ongoing security validation
5. **Audit Trail**: Complete logging of MCP activities

### 🛡️ Security Components

#### 1. Security Assessment Framework
- **Location**: `bin/executable_mcp-security-check`
- **Purpose**: Pre-installation security evaluation
- **Checks**:
  - Repository trust indicators (stars, forks, maintenance)
  - Package vulnerability scanning
  - Network requirements analysis
  - Data access scope evaluation

#### 2. Configuration Security
- **Location**: `dot_claude/settings.json`
- **Features**:
  - Human approval required for all operations
  - Resource indicators (RFC 8707) for OAuth
  - Network access restrictions
  - Command execution blacklist
  - Comprehensive audit logging

#### 3. Environment Security
- **Location**: `private_dot_env.tmpl`
- **Features**:
  - Secure token management
  - Environment-specific configuration
  - Token rotation scheduling
  - Security parameter enforcement

## Installed MCP Servers

### ✅ Context7 (Documentation Server)
- **Risk Level**: LOW
- **Provider**: Upstash (reputable company)
- **Access**: Read-only documentation fetching
- **Network**: External documentation sources
- **Security Measures**:
  - Local npm package installation
  - Minimal permissions
  - Regular update monitoring

### ✅ Notion (Productivity Server)
- **Risk Level**: MEDIUM
- **Provider**: Official Notion implementation
- **Access**: Full workspace access (user account level)
- **Network**: api.notion.com
- **Security Measures**:
  - OAuth bearer token authentication
  - Resource indicators implementation
  - Human approval for all operations
  - Regular token rotation

## Security Procedures

### 🔍 Pre-Installation Assessment

Before installing any MCP server:

1. **Run Security Assessment**:
   ```bash
   mcp-security-check <server_name> --repo-url <url> --package-name <package> --data-scope <scope>
   ```

2. **Review Assessment Results**:
   - ✅ Approved: Safe to install
   - ⚠️ Caution: Install with enhanced monitoring
   - ❌ Rejected: Manual review required

3. **Validate Configuration**:
   - Check security settings in Claude config
   - Verify environment variables
   - Test audit logging

### 🔧 Installation Process

1. **Automated Setup**: Use `run_once_05-setup-secure-mcp.sh.tmpl`
2. **Manual Verification**: Check `claude mcp list` for proper configuration
3. **Authentication Setup**: Configure required tokens in `.env`
4. **Security Validation**: Review audit logs after first use

### 📊 Ongoing Monitoring

#### Daily
- Review audit logs: `~/.local/share/claude/audit.log`
- Monitor for unusual MCP activity

#### Weekly
- Update MCP packages: `npm update -g @upstash/context7-mcp`
- Check for security advisories
- Rotate authentication tokens (if required)

#### Monthly
- Full security assessment of installed servers
- Review and update security configurations
- Archive old audit logs

## Authentication & Token Management

### 🔑 Required Tokens

1. **Notion API Token**:
   ```bash
   # Generate at: https://www.notion.so/my-integrations
   export NOTION_API_TOKEN="secret_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
   ```

2. **GitHub Token** (for future GitHub MCP):
   ```bash
   # Generate at: https://github.com/settings/tokens
   export GITHUB_TOKEN="ghp_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
   ```

### 🔄 Token Rotation

- **Frequency**: Every 30 days (configurable)
- **Process**:
  1. Generate new token
  2. Update environment variable
  3. Test MCP server connectivity
  4. Revoke old token
  5. Update audit log

## Incident Response

### 🚨 Security Incident Types

1. **Unauthorized MCP Access**: Suspicious commands in audit log
2. **Token Compromise**: Unexpected API usage
3. **Malicious MCP Server**: Suspicious behavior from installed server

### 📋 Response Procedures

1. **Immediate**:
   ```bash
   # Disable all MCP servers
   claude mcp list | grep -v "No MCP" | while read server; do
     claude mcp remove "$(echo $server | cut -d: -f1)"
   done
   ```

2. **Investigation**:
   - Review audit logs
   - Check network activity
   - Validate installed packages

3. **Recovery**:
   - Rotate all authentication tokens
   - Reinstall MCP servers after security review
   - Update security configurations

## Configuration Files

### 📁 File Locations

- **Security Script**: `bin/executable_mcp-security-check`
- **Claude Config**: `dot_claude/settings.json`
- **Environment**: `private_dot_env.tmpl`
- **Setup Script**: `run_once_05-setup-secure-mcp.sh.tmpl`
- **Audit Logs**: `~/.local/share/claude/audit.log`

### 🔧 Key Settings

```json
{
  "security": {
    "requireHumanApproval": true,
    "enableResourceIndicators": true,
    "maxExecutionTime": 30000,
    "allowedNetworkTargets": [
      "api.notion.com",
      "api.github.com",
      "raw.githubusercontent.com"
    ],
    "blockedCommands": [
      "rm -rf",
      "sudo",
      "chmod 777"
    ]
  }
}
```

## Best Practices

### ✅ Do's

- Always run security assessment before installing MCP servers
- Keep packages updated regularly
- Monitor audit logs frequently
- Use environment-specific configurations
- Implement token rotation schedules
- Test MCP servers in isolated environments first

### ❌ Don'ts

- Don't install MCP servers without security review
- Don't store tokens in plain text
- Don't ignore security warnings
- Don't grant unnecessary permissions
- Don't skip audit log reviews
- Don't use MCP servers from unknown sources

## Troubleshooting

### 🔧 Common Issues

1. **MCP Server Won't Connect**:
   - Check package installation
   - Verify command syntax
   - Review authentication tokens

2. **Security Assessment Fails**:
   - Update security assessment script
   - Check GitHub API access
   - Review package registry access

3. **Authentication Errors**:
   - Verify token format
   - Check token permissions
   - Ensure token isn't expired

### 📞 Support Resources

- **Claude Code MCP Documentation**: https://docs.anthropic.com/en/docs/claude-code/mcp
- **Security Issues**: Review this documentation and audit logs
- **Package Issues**: Check npm registry and package documentation

## Updates & Maintenance

This security framework is designed to evolve with new threats and MCP developments. Regular updates include:

- Security assessment script improvements
- New security checks for emerging threats
- Updated best practices and procedures
- Enhanced monitoring and alerting capabilities

---

**Last Updated**: $(date +%Y-%m-%d)
**Version**: 1.0.0
**Maintainer**: Automated via chezmoi dotfiles