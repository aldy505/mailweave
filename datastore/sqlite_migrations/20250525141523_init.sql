-- +goose Up
-- +goose StatementBegin
CREATE TABLE mailweave_tls_rpt_report (
    id INTEGER PRIMARY KEY,
    raw_report TEXT NOT NULL,
    domain_owner TEXT,
    organization_name TEXT,
    domain_name TEXT,
    report_id TEXT,
    extra_contact_info TEXT,
    report_date TEXT,
    range_start TEXT,
    range_end TEXT,
    email_sender TEXT,
    email_recipient TEXT,
    email_subject TEXT,
    report_file_name TEXT,
    total_sessions INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE mailweave_tls_rpt_report_row (
    id INTEGER PRIMARY KEY,
    report_id INTEGER NOT NULL,
    domain_name TEXT NOT NULL,
    ip_address TEXT,
    policy_type TEXT,
    policy_string TEXT,
    mx_host TEXT,
    successful_count INTEGER,
    failed_count INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (report_id) REFERENCES mailweave_tls_rpt_report(id) ON DELETE CASCADE
);

CREATE TABLE mailweave_tls_rpt_aggregate (
    id INTEGER PRIMARY KEY,
    domain_owner TEXT NOT NULL,
    organization_name TEXT,
    domain TEXT NOT NULL,
    successful_percentage FLOAT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (domain_owner, organization_name, domain)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE mailweave_tls_rpt_report;
DROP TABLE mailweave_tls_rpt_report_row;
DROP TABLE mailweave_tls_rpt_aggregate;
-- +goose StatementEnd
