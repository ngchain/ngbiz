package ngsheet

import (
	"bytes"
	"github.com/ngchain/ngcore/ngtypes"
)

func (m *Manager) AccountIsRegistered(accountID uint64) bool {
	m.accountsMu.RLock()
	defer m.accountsMu.RUnlock()

	_, exists := m.accounts[accountID]
	return exists
}

func (m *Manager) GetAccountById(id uint64) (account *ngtypes.Account, exists bool) {
	m.accountsMu.RLock()
	defer m.accountsMu.RUnlock()

	account, exists = m.accounts[id]

	return
}

func (m *Manager) GetAccountsByPublicKey(publicKey []byte) []*ngtypes.Account {
	m.accountsMu.RLock()
	defer m.accountsMu.RUnlock()

	accounts := make([]*ngtypes.Account, 0)
	for _, account := range m.accounts {
		if bytes.Compare(account.Owner, publicKey) == 0 {
			accounts = append(accounts, account)
		}
	}

	return accounts
}

// RegisterAccount is same to balanceSheet RegisterAccount, this is for consensus calling
func (m *Manager) RegisterAccount(account *ngtypes.Account) (ok bool) {
	m.accountsMu.Lock()
	defer m.accountsMu.Unlock()

	if _, exists := m.accounts[account.ID]; !exists {
		m.accounts[account.ID] = account
		return true
	}

	return false
}

func (m *Manager) DeleteAccount(account *ngtypes.Account) (ok bool) {
	m.accountsMu.Lock()
	defer m.accountsMu.Unlock()

	if _, exists := m.accounts[account.ID]; !exists {
		return false
	}

	delete(m.accounts, account.ID)
	return true
}