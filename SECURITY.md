# Security Guide

## Handling Sensitive Configuration

### Current Status
✅ `conf/deploy.yaml` has been removed from git tracking
✅ Added to `.gitignore` to prevent future commits
✅ Created `conf/deploy.yaml.example` as a template

### Removing Sensitive Data from Git History

The sensitive configuration file was committed in earlier commits. To completely remove it from git history, you have two options:

#### Option 1: Using git-filter-repo (Recommended)

```bash
# Install git-filter-repo if not already installed
pip3 install git-filter-repo

# Remove the file from all history
git filter-repo --path conf/deploy.yaml --invert-paths

# Force push (WARNING: This rewrites history)
git push origin --force --all
```

#### Option 2: Using BFG Repo-Cleaner

```bash
# Download BFG from https://rtyley.github.io/bfg-repo-cleaner/
# Remove the file
java -jar bfg.jar --delete-files deploy.yaml

# Clean up
git reflog expire --expire=now --all
git gc --prune=now --aggressive

# Force push
git push origin --force --all
```

⚠️ **Warning**: Force pushing rewrites git history. If others have cloned the repository, they'll need to re-clone it.

### Best Practices for Production

1. **Use Google Secret Manager** (Recommended for GAE):
   ```go
   // Access secrets from Secret Manager instead of files
   secret, err := secretmanager.AccessSecretVersion(ctx, "projects/PROJECT_ID/secrets/SECRET_NAME/versions/latest")
   ```

2. **Use Environment Variables**:
   ```yaml
   # In app.yaml
   env_variables:
     ELASTICSEARCH_ADDRESS: "http://..."
     ELASTICSEARCH_USERNAME: "username"
     ELASTICSEARCH_PASSWORD: "password"
   ```

3. **Use .env files** (for local development only):
   - Never commit `.env` files
   - Add `.env` to `.gitignore`
   - Use libraries like `godotenv` to load them

4. **Rotate Credentials**:
   - After removing from git, consider rotating all exposed credentials
   - Update passwords and secrets that were in the repository

### What to Do Now

1. ✅ Configuration is now excluded from git
2. ⚠️ Consider rotating exposed credentials (Elasticsearch password, JWT secret)
3. ⚠️ Remove sensitive data from git history (see options above)
4. ✅ Use Secret Manager or environment variables for production
